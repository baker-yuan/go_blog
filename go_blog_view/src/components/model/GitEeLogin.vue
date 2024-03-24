<template>
  <div>
    <div>gitee登陆中...</div>
  </div>
</template>

<script>
export default {
  data() {
    return {}
  },
  mounted() {
    const that = this
    const code = this.$route.query.code;
    this.axios.post("/user/login/giteeOauth", {
      'code': code
    }).then(({data}) => {
      if (data && data.status === 200) {
        that.$store.commit("login", data.data);
        that.$store.commit("closeModel");
        that.$toast({ type: "success", message: "登录成功" });
        this.$router.push({name:'user'})

      } else {
        that.$toast({ type: "error", message: data.desc });
      }
    });
  }
};
</script>

<style scoped>

</style>