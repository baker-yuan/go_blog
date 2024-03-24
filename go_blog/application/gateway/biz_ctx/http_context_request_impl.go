package biz_ctx

import (
	"bytes"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
)

// URIRequest url查询参数读取
type URIRequest struct {
	uri *fasthttp.URI
}

// 强制URIRequest实现IURIReader
var _ IURIReader = (*URIRequest)(nil)

// 强制URIRequest实现IURIWriter
var _ IURIWriter = (*URIRequest)(nil)

func NewURIRequest(uri *fasthttp.URI) *URIRequest {
	return &URIRequest{uri: uri}
}

func (ur *URIRequest) GetQuery(key string) string {
	return string(ur.uri.QueryArgs().Peek(key))
}

func (ur *URIRequest) RawQuery() string {
	return string(ur.uri.QueryString())
}

func (ur *URIRequest) RequestURI() string {
	return string(ur.uri.RequestURI())
}

func (ur *URIRequest) Scheme() string {
	return string(ur.uri.Scheme())
}

func (ur *URIRequest) RawURL() string {
	return string(ur.uri.FullURI())
}

func (ur *URIRequest) Host() string {
	return string(ur.uri.Host())
}

func (ur *URIRequest) Path() string {
	return string(ur.uri.Path())
}

func (ur *URIRequest) reset(uri *fasthttp.URI) {
	ur.uri = uri
}

func (ur *URIRequest) SetQuery(key, value string) {
	ur.uri.QueryArgs().Set(key, value)
}

func (ur *URIRequest) AddQuery(key, value string) {
	ur.uri.QueryArgs().Add(key, value)
}

func (ur *URIRequest) DelQuery(key string) {
	queryArgs := ur.uri.QueryArgs()
	queryArgs.Del(key)
	if queryArgs.Len() == 0 {
		ur.uri.SetQueryStringBytes(nil)
	}
}

func (ur *URIRequest) SetRawQuery(raw string) {
	ur.uri.SetQueryString(raw)
}

func (ur *URIRequest) SetPath(s string) {
	ur.uri.SetPath(s)

}

func (ur *URIRequest) SetScheme(scheme string) {
	ur.uri.SetScheme(scheme)
}

func (ur *URIRequest) SetHost(host string) {
	ur.uri.SetHost(host)
}

// RequestHeader 请求头读取
type RequestHeader struct {
	header *fasthttp.RequestHeader // 请求头读取
	tmp    http.Header             // map[string][]string
}

// 强制RequestHeader实现IHeaderWriter
var _ IHeaderWriter = (*RequestHeader)(nil)

// 强制RequestHeader实现IHeaderReader
var _ IHeaderReader = (*RequestHeader)(nil)

// reset 初始化
func (h *RequestHeader) reset(header *fasthttp.RequestHeader) {
	h.header = header
	h.tmp = nil
}

func (h *RequestHeader) RawHeader() string {
	return h.header.String()
}

func (h *RequestHeader) GetHeader(name string) string {
	return h.Headers().Get(name)
}

func (h *RequestHeader) Headers() http.Header {
	h.initHeader()
	return h.tmp
}

func (h *RequestHeader) Host() string {
	return string(h.header.Host())
}

func (h *RequestHeader) GetCookie(key string) string {
	return string(h.header.Cookie(key))
}

func (h *RequestHeader) initHeader() {
	if h.tmp == nil {
		h.tmp = make(http.Header)
		h.header.VisitAll(func(key, value []byte) {
			bytes.SplitN(value, []byte(":"), 2)
			h.tmp[string(key)] = []string{string(value)}
		})
	}
}

func (h *RequestHeader) SetHeader(key, value string) {
	if h.tmp != nil {
		h.tmp.Set(key, value)
	}
	h.header.Set(key, value)
}

func (h *RequestHeader) AddHeader(key, value string) {
	if h.tmp != nil {
		h.tmp.Add(key, value)
	}
	h.header.Add(key, value)
}

func (h *RequestHeader) DelHeader(key string) {
	if h.tmp != nil {
		h.tmp.Del(key)
	}
	h.header.Del(key)
}

func (h *RequestHeader) SetHost(host string) {
	if h.tmp != nil {
		h.tmp.Set("Host", host)
	}
	h.header.SetHost(host)
}

// HTTP协议中的内容类型（Content-Type）
const (
	// MultipartForm
	// 主要用于在表单中发送二进制数据。最常见的用例是上传文件。在此编码类型中，表单的每个字段被视为一部分（multipart），每个部分都包含有关该字段的信息，例如字段名，字段值，如果字段是文件，则还包含文件名和文件类型。这意味着，使用这种类型，你可以在同一请求中发送文本和数据。
	MultipartForm = "multipart/form-data"
	// FormData
	// 通常用于发送ASCII字符集的数据。在此编码类型中，表单的字段名和值用等号（=）连接，字段之间用&符号分隔。所有非字母数字字符都会被百分比编码。这种类型常用于提交简单的文本数据。
	FormData = "application/x-www-form-urlencoded"
	// TEXT 表示纯文本数据。
	TEXT = "text/plain"
	// JSON 表示 JSON 格式的数据，常用于客户端和服务器之间的数据交换。
	JSON           = "application/json"
	JavaScript     = "application/javascript"
	AppLicationXML = "application/xml"
	TextXML        = "text/xml"
	Html           = "text/html"
)

var (
	ErrorNotForm      = errors.New("contentType is not Form")
	ErrorNotMultipart = errors.New("contentType is not Multipart")
	ErrorNotAllowRaw  = errors.New("contentType is not allow Raw")
	ErrorNotSend      = errors.New("not send")
)

// BodyRequestHandler 请求体读取
type BodyRequestHandler struct {
	request  *fasthttp.Request
	formdata *multipart.Form
}

// 强制IBodyDataWriter实现IHeaderReader
var _ IBodyDataWriter = (*BodyRequestHandler)(nil)

// 强制IBodyDataWriter实现IBodyDataReader
var _ IBodyDataReader = (*BodyRequestHandler)(nil)

func NewBodyRequestHandler(request *fasthttp.Request) *BodyRequestHandler {
	return &BodyRequestHandler{request: request}
}

func (b *BodyRequestHandler) reset(request *fasthttp.Request) {
	b.request = request
	b.formdata = nil
}

// ContentType 获取contentType
func (b *BodyRequestHandler) ContentType() string {
	return string(b.request.Header.ContentType())
}

// BodyForm 获取表单参数
func (b *BodyRequestHandler) BodyForm() (url.Values, error) {
	contentType, _, _ := mime.ParseMediaType(string(b.request.Header.ContentType()))
	switch contentType {
	case FormData:
		return url.ParseQuery(string(b.request.Body()))
	case MultipartForm:
		multipartForm, err := b.MultipartForm()
		if err != nil {
			return nil, err
		}
		return multipartForm.Value, nil
	default:
		return nil, ErrorNotForm
	}

}

func (b *BodyRequestHandler) Files() (map[string][]*multipart.FileHeader, error) {
	form, err := b.MultipartForm()
	if err != nil {
		return nil, err
	}
	return form.File, nil
}

// GetForm 获取表单参数
func (b *BodyRequestHandler) GetForm(key string) string {
	contentType, _, _ := mime.ParseMediaType(b.ContentType())
	switch contentType {
	case FormData:
		args := b.request.PostArgs()
		if args == nil {
			return ""
		}
		return string(args.Peek(key))
	case MultipartForm:
		form, err := b.MultipartForm()
		if err != nil {
			return ""
		}
		vs := form.Value[key]
		if len(vs) > 0 {
			return vs[0]
		}
		return ""
	}
	return ""
}

func (b *BodyRequestHandler) GetFile(key string) ([]*multipart.FileHeader, bool) {
	multipartForm, err := b.MultipartForm()
	if err != nil {
		return nil, false
	}
	fl, has := multipartForm.File[key]
	return fl, has
}

// RawBody 获取raw数据
func (b *BodyRequestHandler) RawBody() ([]byte, error) {
	return b.request.Body(), nil
}

func (b *BodyRequestHandler) MultipartForm() (*multipart.Form, error) {
	if b.formdata != nil {
		return b.formdata, nil
	}
	if !strings.Contains(b.ContentType(), MultipartForm) {
		return nil, ErrorNotMultipart
	}
	form, err := b.request.MultipartForm()
	if err != nil {
		return nil, err
	}

	b.formdata = &multipart.Form{
		Value: form.Value,
		File:  form.File,
	}
	_ = b.resetFile()
	return form, nil
}

func (b *BodyRequestHandler) resetFile() error {
	multipartForm := b.formdata
	if multipartForm == nil {
		return nil
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for name, fs := range multipartForm.File {
		for _, f := range fs {
			fio, err := f.Open()
			if err != nil {
				return err
			}

			part, err := writer.CreateFormFile(name, f.Filename)
			if err != nil {
				_ = fio.Close()
				return err
			}

			data, err := io.ReadAll(fio)
			if err != nil {
				_ = fio.Close()
				return err
			}
			_, err = part.Write(data)
			if err != nil {
				_ = fio.Close()
				return err
			}
			_ = fio.Close()
		}
	}

	for key, values := range multipartForm.Value {
		for _, value := range values {
			err := writer.WriteField(key, value)
			if err != nil {
				return err
			}
		}

	}
	err := writer.Close()
	if err != nil {
		return err
	}
	b.request.Header.SetContentType(writer.FormDataContentType())
	b.request.SetBodyRaw(body.Bytes())
	return nil
}

// SetForm 设置表单参数
func (b *BodyRequestHandler) SetForm(values url.Values) error {
	contentType, _, _ := mime.ParseMediaType(b.ContentType())
	if contentType != FormData && contentType != MultipartForm {
		return ErrorNotForm
	}
	switch contentType {
	case FormData:
		b.request.SetBodyString(values.Encode())
	case MultipartForm:
		multipartForm, err := b.MultipartForm()
		if err != nil {
			return err
		}
		multipartForm.Value = values
		return b.resetFile()
	}
	return ErrorNotForm
}

func (b *BodyRequestHandler) SetToForm(key, value string) error {
	contentType, _, _ := mime.ParseMediaType(string(b.request.Header.ContentType()))
	switch contentType {
	case FormData:
		b.request.PostArgs().Set(key, value)
		b.request.SetBodyRaw(b.request.PostArgs().QueryString())
		return nil
	case MultipartForm:
		multipartForm, err := b.MultipartForm()
		if err != nil {
			return err
		}
		multipartForm.Value[key] = []string{value}
		return b.resetFile()
	default:
		return ErrorNotForm
	}
}

// AddForm 新增表单参数
func (b *BodyRequestHandler) AddForm(key, value string) error {
	contentType, _, _ := mime.ParseMediaType(string(b.request.Header.ContentType()))
	switch contentType {
	case FormData:
		b.request.PostArgs().Add(key, value)
		b.request.SetBody(b.request.PostArgs().QueryString())
		return nil
	case MultipartForm:
		multipartForm, err := b.MultipartForm()
		if err != nil {
			return err
		}
		multipartForm.Value[key] = append(multipartForm.Value[key], value)
		return b.resetFile()
	default:
		return ErrorNotForm
	}
}

// AddFile 新增文件参数
func (b *BodyRequestHandler) AddFile(key string, file *multipart.FileHeader) error {
	contentType, _, _ := mime.ParseMediaType(b.ContentType())
	if contentType != FormData && contentType != MultipartForm {
		return ErrorNotMultipart
	}
	multipartForm, err := b.MultipartForm()
	if err != nil {
		return err
	}
	multipartForm.File[key] = append(multipartForm.File[key], file)
	return b.resetFile()
}

// SetRaw 设置raw数据
func (b *BodyRequestHandler) SetRaw(contentType string, body []byte) {
	b.request.SetBodyRaw(body)
	b.request.Header.SetContentType(contentType)
}

// RequestReader 请求数据读取接口
type RequestReader struct {
	req        *fasthttp.Request  // 客户端和网关之间的请求(拷贝后的请求)
	body       BodyRequestHandler // 请求体读取
	headers    RequestHeader      // 请求头读取
	uri        URIRequest         // url查询参数读取
	remoteAddr string
	remotePort string
	realIP     string
	length     int
}

// 强制RequestReader实现IRequestReader
var _ IRequestReader = (*RequestReader)(nil)

// reset 重置
func (r *RequestReader) reset(req *fasthttp.Request, remoteAddr string) {
	r.req = req
	r.remoteAddr = remoteAddr
	r.body.reset(req)
	r.headers.reset(&req.Header)
	r.uri.uri = req.URI()

	idx := strings.LastIndex(remoteAddr, ":")
	if idx != -1 {
		r.remoteAddr = remoteAddr[:idx]
		r.remotePort = remoteAddr[idx+1:]
	}
	length := r.req.Header.ContentLength()
	if length > 0 {
		r.length = length
	}
}

func (r *RequestReader) URI() IURIReader {
	return &r.uri
}

func (r *RequestReader) Header() IHeaderReader {
	return &r.headers
}

func (r *RequestReader) Body() IBodyDataReader {
	return &r.body
}

func (r *RequestReader) RemoteAddr() string {
	return r.remoteAddr
}

func (r *RequestReader) RemotePort() string {
	return r.remotePort
}

func (r *RequestReader) RealIP() string {
	if r.realIP == "" {
		realIP := r.headers.GetHeader("x-real-ip")
		if realIP == "" {
			realIP = r.remoteAddr
		}
		r.realIP = realIP
	}
	return r.realIP
}

func (r *RequestReader) ForwardIP() string {
	return r.headers.GetHeader("x-forwarded-for")
}

func (r *RequestReader) Method() string {
	return string(r.req.Header.Method())
}

func (r *RequestReader) ContentLength() int {
	return r.length
}

func (r *RequestReader) ContentType() string {
	return string(r.req.Header.ContentType())
}

func (r *RequestReader) String() string {
	return r.req.String()
}

func (r *RequestReader) Request() *fasthttp.Request {
	return r.req
}

func (r *RequestReader) Finish() error {
	r.req = nil
	r.body.reset(nil)
	r.headers.reset(nil)
	r.uri.reset(nil)
	return nil
}
