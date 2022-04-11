<template>
  <div
    class="icon-container"
    :class="{ 'rounded-button': type === 'rounded' }"
    :style="{ 'height': size + 'px', 'width': size + 'px', 'padding': padding + 'px', 'background-color': (isHover && hover) ? 'var(--complementary-white-5)' : bgColor, 'border-radius': brdRadius + 'px' }"
    @mouseover="isHover = true"
    @mouseleave="isHover = false"
  >
    <img :style="{ 'height': size + 'px', 'width': size + 'px' }" :src="require(`@/assets/img/icons/${name}.svg`)" :alt="alt">
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
    },
    padding: {
      type: String,
      required: false,
      default: '0'
    },
    bgColor: {
      type: String,
      required: false,
      default: 'rgba(255, 255, 255, 0)'
    },
    hover: {
      type: Boolean,
      required: false,
      default: false
    },
    brdRadius: {
      type: String,
      required: false,
      default: '0'
    }
  },
  data () {
    return {
      isHover: false
    }
  }
}
</script>

<style scoped>
  .icon-container {
    position: relative;
  }

  .icon-container img {
    position: absolute;
  }

  .rounded-button {
    padding: 5px;
    border-radius: 100%;
  }
</style>
