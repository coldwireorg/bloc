<template>
  <div>
    <ColdwireLogo />
    <AppSideBarMobile />
    <div class="appcontent-rsp-container">
      <Nuxt />
    </div>
  </div>
</template>

<script>
export default {
  name: 'MobileLayout'
}
</script>

<style scoped>
  .appcontent-rsp-container {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 106px;
    background-color: #161B22;
  }
</style>
