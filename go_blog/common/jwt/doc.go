// Package jwt 提供了用于创建和验证 JSON Web Tokens (JWT) 的工具。
// 它支持使用 RSA 签名来生成和验证令牌，并允许用户自定义密钥对。
//
// 主要组件包括：
// - TokenGeneratorConfig：用于生成JWT的配置，包含私钥和令牌有效期等信息。
// - TokenValidatorConfig：用于验证JWT的配置，包含公钥信息。
// - TokenGenerator：负责生成带有 RSA 签名的访问令牌和刷新令牌。
// - TokenValidator：负责验证令牌的签名并提取令牌中的用户信息。
// - UserInfo：定义了存储在 JWT 中的用户信息结构。
// - TokenDetails：定义了访问令牌和刷新令牌的详细信息结构。
//
// 使用方法：
// 生成Token:
//  1. 创建TokenGeneratorConfig实例并填充必要的配置信息。建议直接用TokenGeneratorConfig去映射配置文件。
//  2. 调用TokenGeneratorConfig的Build方法创建TokenGenerator实例。
//  3. 使用TokenGenerator的CreateToken方法生成访问令牌和刷新令牌。
//
// 使用Token:
//  1. 创建TokenValidatorConfig实例并填充公钥信息。。建议直接用TokenGeneratorConfig去映射配置文件。
//  2. 调用TokenValidatorConfig的Build方法创建TokenValidator实例。
//  3. 使用TokenValidator的ValidateAndExtractUserInfo方法验证访问令牌并提取用户信息。
//
// 工具函数：
//   - GenerateBase64EncodedKeys：生成Base64编码的RSA密钥对，用于签名和验证JWT。用户可以使用此函数方便地生成新的密钥对。
//
// 此外，包中还包含了一系列错误定义，用于处理各种 JWT 相关的异常情况。
// utils.go 提供了用于解析 Base64 编码的 RSA 密钥的辅助函数。
//
// 该包旨在为用户服务和网关服务提供JWT相关的安全和验证机制。
package jwt
