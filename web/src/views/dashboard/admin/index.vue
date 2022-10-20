<template>
  <div class="dashboard-editor-container">
    <h1 style="font-size:2em;margin:.67em 0">欢迎回来,{{ userName }}</h1>
  </div>
</template>

<script>
import { getInfo } from '@/api/user'

export default {
  name: 'DashboardAdmin',
  components: {},
  data() {
    return {
      userName: ''
    }
  },
  created() {
    this.getUserInfo()
  },
  methods: {
    getUserInfo() {
      getInfo().then((response) => {
        console.log(response)
        this.userName = response.data.userName
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 12px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

/deep/ .el-tabs__item {
  padding-left: 16px !important;
  height: 50px;
  line-height: 50px;
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
