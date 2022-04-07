<template>
  <div>
    <a v-if="isExternal" target="_blank" :href="link" class="action-button" :class="getClassColor()">
      {{ text }}
      <slot />
    </a>
    <nuxt-link v-else :to="link" class="action-button" :class="getClassColor()">
      {{ text }}
      <slot />
    </nuxt-link>
  </div>
</template>

<script>
export default {
  name: 'ButtonAction',
  props: {
    text: {
      type: String,
      required: true
    },
    link: {
      type: String,
      required: false,
      default: '#'
    },
    isExternal: {
      type: Boolean,
      required: false,
      default: false
    },
    color: {
      type: String,
      required: false,
      default: 'green'
    }
  },

  methods: {
    getClassColor () {
      if (this.color === 'blue' || this.color === 'BLUE') {
        return 'prm-blue'
      } else {
        return 'prm-green'
      }
    }
  }
}
</script>

<style scoped>
  .action-button {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    box-sizing: border-box;
    padding: 8px 9px;
    gap: 6px;
    border-radius: 8px;
  }

  .prm-green {
    background: var(--primary-green);
  }

  .prm-blue {
    background: var(--primary-blue);
  }
</style>
