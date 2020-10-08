<template>
  <input
    :id="id"
    :value="value"
    :type="type"
    class="f6 input-reset ba br2 b--black-10 pa2 w5"
    @input="$emit('input', $event.target.value)"
  />
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component
export default class extends Vue {
  @Prop({ type: String, required: true })
  private readonly id!: string

  @Prop({ type: String, required: true })
  private readonly type!: string

  @Prop({ type: String, required: true })
  private readonly value!: string
}
</script>
