import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    //


    searchFlag: false,
    loginFlag: false,
    registerFlag: false,
    forgetFlag: false,
    emailFlag: false,

    drawer: false,
    loginUrl: "",

    uid: null,
    avatar: null,
    nickname: null,
    intro: null,
    webSite: null,
    loginType: null,
    email: null,

    // todo 废弃 articleLikeSet commentLikeSet
    commentLikeSet: [],
    // 博客配置信息
    blogInfo: {},
    // oauth登陆信息
    oauthLoginConfig: {}
  },
  mutations: {
    login(state, user) {
      window.localStorage.setItem("token", user.ticket);
      state.uid = user.userInfo.uid;
      state.avatar = user.userInfo.avatar;
      state.nickname = user.userInfo.nickname;
      state.intro = user.userInfo.intro;
      state.webSite = user.userInfo.webSite;
      state.email = user.userInfo.email;
      state.loginType = user.userInfo.loginType;
    },
    logout(state) {
      window.localStorage.setItem("token", "");
      state.uid = null;
      state.avatar = null;
      state.nickname = null;
      state.intro = null;
      state.webSite = null;
      state.email = null;
      state.loginType = null;
    },
    saveLoginUrl(state, url) {
      state.loginUrl = url;
    },
    saveEmail(state, email) {
      state.email = email;
    },
    updateUserInfo(state, user) {
      state.nickname = user.nickname;
      state.intro = user.intro;
      state.webSite = user.webSite;
    },
    savePageInfo(state, pageList) {
      state.pageList = pageList;
    },
    saveOauthLoginConfig(state, oauthLoginConfig) {
      state.oauthLoginConfig = oauthLoginConfig;
    },
    updateAvatar(state, avatar) {
      state.avatar = avatar;
    },
    checkBlogInfo(state, blogInfo) {
      state.blogInfo = blogInfo;
    },
    closeModel(state) {
      state.registerFlag = false;
      state.loginFlag = false;
      state.searchFlag = false;
      state.emailFlag = false;
    },
    commentLike(state, commentId) {
      var commentLikeSet = state.commentLikeSet;
      if (commentLikeSet.indexOf(commentId) != -1) {
        commentLikeSet.splice(commentLikeSet.indexOf(commentId), 1);
      } else {
        commentLikeSet.push(commentId);
      }
    }
  },
  actions: {},
  modules: {},
  plugins: [
    createPersistedState({
      storage: window.sessionStorage
    })
  ]
});
