<template>
  <div>
    <img v-if="name === 'disconnect'" src="@/assets/img/icons/disconnect16x16.svg" :alt="alt">
    <img v-else-if="name == 'folder'" src="@/assets/img/icons/folder16x16.svg" :alt="alt">
    <img v-else-if="name == 'heart-around'" src="@/assets/img/icons/heart-around16x16.svg" :alt="alt">
    <img v-else-if="name == 'share'" src="@/assets/img/icons/share16x16.svg" :alt="alt">
    <img v-else-if="name == 'star'" src="@/assets/img/icons/star16x16.svg" :alt="alt">
    <img v-else-if="name == 'upload'" src="@/assets/img/icons/upload16x16.svg" :alt="alt">
    <img v-else src="@/assets/img/icons/disconnect16x16.svg" :alt="alt">
  </div>
</template>

<script>
export default {
  name: 'IconComponent',
  props: {
    name: {
      type: String,
      required: false,
      default: 'unknow'
    },
    alt: {
      type: String,
      required: false,
      default: ''
    }
  }
}
</script>
