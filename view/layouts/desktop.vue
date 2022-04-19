<template>
  <div>
    <ColdwireLogo />
    <AppSideBarDesktop />
    <div class="appcontent-rsp-container">
      <Nuxt />
    </div>
  </div>
</template>

<script>
export default {
  name: 'DesktopLayout'
}
</script>

<style scoped>
  .appcontent-rsp-container {
    position: absolute;
    top: 0;
    left: 254px;
    right: 0;
    bottom: 0;
    background-color: #161B22;
  }
</style>
