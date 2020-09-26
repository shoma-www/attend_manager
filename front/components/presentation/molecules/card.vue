<template>
  <div class="ba br3 shadow-5">
    <h1 class="f4 bg-black-70 br3 br--top white mv0 pv2 ph3">{{ title }}</h1>
    <div class="pa3">
      <slot />
    </div>
  </div>
</template>

<script lang="ts">
import { Prop } from 'nuxt-property-decorator'
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {
  @Prop({ type: String, required: true })
  private readonly title!: string
}
</script>
