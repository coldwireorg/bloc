<template>
  <div />
</template>

<script>
export default {
  name: 'AppPage',
  layout (context) {
    return context.$device.isMobile ? 'mobile' : 'desktop'
  }
}
</script>
