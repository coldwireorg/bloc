<template>
  <div>
    <ColdwireLogo />
    <AppSideBarDesktop />
    <div class="appcontent-position">
      <div class="fx-column appcontent-main">
        <Nuxt />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DesktopLayout'
}
</script>

<style scoped>
  .appcontent-position {
    position: absolute;
    top: 0;
    left: 254px;
    right: 0;
    bottom: 0;
    background-color: #161B22;
  }

  .appcontent-main {
    justify-content: flex-start;
    flex-wrap: nowrap;
    padding: 1rem 2rem;
  }
</style>
