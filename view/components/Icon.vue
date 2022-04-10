<template>
  <div
    class="icon-container"
    :class="{
      'icon-size-16': size === '16',
      'icon-size-24': size === '24',
      'rounded-button': type === 'rounded'
    }"
  >
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
    },
    size: {
      type: String,
      required: false,
      default: '16'
    },
    type: {
      type: String,
      required: false,
      default: 'normal'
    }
  }
}
</script>

<style scoped>
  .icon-container {
    position: relative;
    box-sizing: border-box;
  }

  .icon-container img {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
  }

  .rounded-button {
    padding: 5px;
    border-radius: 100%;
  }

  .rounded-button:hover {
    background-color: var(--complementary-white-5);
  }

  .icon-size-16 {
    height: 16px;
    width: 16px;
  }
  .icon-size-24 {
    height: 24px;
    width: 24px;
  }
</style>
