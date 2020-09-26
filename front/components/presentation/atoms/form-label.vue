<template>
  <label :for="id" class="f5 db mb2">{{ label }}</label>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'

@Component
export default class extends Vue {
  @Prop({ type: String, required: true })
  private readonly id!: string

  @Prop({ type: String, required: true })
  private readonly label!: string
}
</script>
