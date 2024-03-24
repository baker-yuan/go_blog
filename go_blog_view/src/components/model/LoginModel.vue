<template>
  <v-dialog v-model="loginFlag" :fullscreen="isMobile" max-width="460">
    <v-card class="login-container" style="border-radius:4px">
      <v-icon class="float-right" @click="loginFlag = false">
        mdi-close
      </v-icon>
      <div class="login-wrapper">
        <!-- 用户名 -->
        <v-text-field
            v-model="loginForm.username"
            label="邮箱号"
            placeholder="请输入您的邮箱号"
            clearable
            @keyup.enter="login"
        />
        <!-- 密码 -->
        <v-text-field
            v-model="loginForm.password"
            class="mt-7"
            label="密码"
            placeholder="请输入您的密码"
            @keyup.enter="login"
            :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
            :type="show ? 'text' : 'password'"
            @click:append="show = !show"
        />
        <!-- 验证码 -->
        <v-row>
          <v-col cols="19">
            <v-text-field
                v-model="loginForm.captcha"
                placeholder="验证码"
            ></v-text-field>
          </v-col>
          <v-col cols="5">
            <img :src="captchaPath" @click="getCaptcha()" alt=""/>
          </v-col>
        </v-row>

        <!-- 按钮 -->
        <v-btn
            class="mt-7"
            block
            color="blue"
            style="color:#fff"
            @click="login"
        >
          登录
        </v-btn>
        <!-- 注册和找回密码 -->
        <div class="mt-10 login-tip">
          <span @click="openRegister">立即注册</span>
          <span @click="openForget" class="float-right">忘记密码?</span>
        </div>
        <div v-if="socialLoginList.length > 0">
          <div class="social-login-title">社交账号登录</div>
          <div class="social-login-wrapper">
            <!-- 微博登录 -->
            <a
                v-if="showLogin('weibo')"
                class="mr-3 iconfont iconweibo"
                style="color:#e05244"
                @click="weiboLogin"
            />
            <!-- qq登录 -->
            <a
                v-if="showLogin('qq')"
                class="mr-3 iconfont iconqq"
                style="color:#00AAEE"
                @click="qqLogin"
            />
            <!-- github -->
            <a
                v-if="showLogin('github')"
                class="mr-3 iconfont icongithub"
                style="color:#000000"
                @click="gitHubLogin"
            />
            <!-- gitee -->
            <a
                v-if="showLogin('gitee')"
                class="iconfont icongitee-fill-round"
                style="color:#c71d24"
                @click="gitEeLogin"
            />
          </div>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>

<script>

import {getUUID} from "../../utils";

export default {
  data: function () {
    return {
      captchaPath: "",
      loginForm: {
        username: "",
        password: "",
        captcha: "",
        uuid: "",
      },
      show: false
    };
  },
  created() {
    this.getCaptcha();
  },
  computed: {
    loginFlag: {
      set(value) {
        this.$store.state.loginFlag = value;
      },
      get() {
        return this.$store.state.loginFlag;
      }
    },
    isMobile() {
      const clientWidth = document.documentElement.clientWidth;
      if (clientWidth > 960) {
        return false;
      }
      return true;
    },
    socialLoginList() {
      return this.$store.state.blogInfo.websiteConfig.socialLoginList;
    },
    showLogin() {
      return function (type) {
        return this.socialLoginList.indexOf(type) !== -1;
      };
    }
  },
  methods: {
    openRegister() {
      this.$store.state.loginFlag = false;
      this.$store.state.registerFlag = true;
    },
    openForget() {
      this.$store.state.loginFlag = false;
      this.$store.state.forgetFlag = true;
    },
    // 获取验证码
    getCaptcha() {
      this.loginForm.uuid = getUUID();
      this.axios
          .get(`/user/login/captcha.jpg?uuid=${this.loginForm.uuid}`)
          .then(({data}) => {
            if (data && data.status === 200) {
              this.captchaPath = data.data;
            }
          });
    },
    // 登陆
    login() {
      const that = this;
      let reg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
      if (!reg.test(this.loginForm.username)) {
        this.$toast({type: "error", message: "邮箱格式不正确"});
        return false;
      }
      if (this.loginForm.password.trim().length === 0) {
        this.$toast({type: "error", message: "密码不能为空"});
        return false;
      }
      //发送登录请求
      that.axios.post("/user/login/byCaptcha", that.loginForm).then(({data}) => {
        if (data && data.status === 200) {
          that.loginForm = {};
          that.$store.commit("login", data.data);
          that.$store.commit("closeModel");
          that.$toast({type: "success", message: "登录成功"});
        } else {
          that.$toast({type: "error", message: data.desc});
        }
      });
    },
    /**
     * github登陆
     */
    gitHubLogin() {
      let oauth_uri = 'https://github.com/login/oauth/authorize'
      let client_id = this.$store.state.oauthLoginConfig.gitHubConfig.gitHubClientId
      let redirect_uri = this.$store.state.oauthLoginConfig.gitHubConfig.gitHubRedirectUri
      window.location.href = `${oauth_uri}?client_id=${client_id}&redirect_url=${redirect_uri}`
    },
    /**
     * gitee登陆
     */
    gitEeLogin() {
      let client_id = this.$store.state.oauthLoginConfig.gitEeConfig.gitEeClientId
      let redirect_uri = this.$store.state.oauthLoginConfig.gitEeConfig.gitEeRedirectUri
      window.location.href = `https://gitee.com/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=code`
    },
    qqLogin() {
      //保留当前路径
      this.$store.commit("saveLoginUrl", this.$route.path);
      if (
          navigator.userAgent.match(
              /(iPhone|iPod|Android|ios|iOS|iPad|Backerry|WebOS|Symbian|Windows Phone|Phone)/i
          )
      ) {
        // eslint-disable-next-line no-undef
        QC.Login.showPopup({
          appId: this.config.QQ_APP_ID,
          redirectURI: this.config.QQ_REDIRECT_URI
        });
      } else {
        window.open(
            "https://graph.qq.com/oauth2.0/show?which=Login&display=pc&client_id=" +
            +this.config.QQ_APP_ID +
            "&response_type=token&scope=all&redirect_uri=" +
            this.config.QQ_REDIRECT_URI,
            "_self"
        );
      }
    },
    weiboLogin() {
      //保留当前路径
      this.$store.commit("saveLoginUrl", this.$route.path);
      window.open(
          "https://api.weibo.com/oauth2/authorize?client_id=" +
          this.config.WEIBO_APP_ID +
          "&response_type=code&redirect_uri=" +
          this.config.WEIBO_REDIRECT_URI,
          "_self"
      );
    }
  }
};
</script>

<style scoped>
.social-login-title {
  margin-top: 1.5rem;
  color: #b5b5b5;
  font-size: 0.75rem;
  text-align: center;
}

.social-login-title::before {
  content: "";
  display: inline-block;
  background-color: #d8d8d8;
  width: 60px;
  height: 1px;
  margin: 0 12px;
  vertical-align: middle;
}

.social-login-title::after {
  content: "";
  display: inline-block;
  background-color: #d8d8d8;
  width: 60px;
  height: 1px;
  margin: 0 12px;
  vertical-align: middle;
}

.social-login-wrapper {
  margin-top: 1rem;
  font-size: 2rem;
  text-align: center;
}

.social-login-wrapper a {
  text-decoration: none;
}
</style>
