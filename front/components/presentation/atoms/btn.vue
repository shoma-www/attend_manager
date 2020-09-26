<template>
  <button
    :type="type"
    :disabled="disabled"
    class="f6 link dim br2 ph3 pv2 mt2 dib white bg-black-70"
  >
    <slot />
  </button>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component
export default class extends Vue {
  @Prop({ type: String, required: true })
  private readonly type!: string

  @Prop({ type: Boolean, required: true })
  private readonly disabled!: boolean
}
</script>
