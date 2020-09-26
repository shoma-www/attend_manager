<template>
  <card title="登録フォーム">
    <form @submit.prevent="submit">
      <div class="mb3">
        <form-label id="user_id" label="User ID" />
        <input-field id="user_id" v-model="userId" type="text" />
      </div>
      <div class="mb3">
        <form-label id="password" label="Password" />
        <input-field id="password" v-model="password" type="password" />
      </div>

      <btn :type="submit">Submit</btn>
    </form>
  </card>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { StatusCodes } from 'http-status-codes'
import formLabel from '../../components/presentation/atoms/form-label.vue'
import inputField from '../../components/presentation/atoms/input-field.vue'
import btn from '../../components/presentation/atoms/btn.vue'
import card from '../../components/presentation/molecules/card.vue'

@Component({
  layout: 'default',
  components: {
    formLabel,
    inputField,
    btn,
    card,
  },
})
export default class extends Vue {
  private userId = ''
  private password = ''

  private async submit(): Promise<void> {
    const res = await this.$axios.post('/user/register', {
      user_id: this.userId,
      password: this.password,
    })
    if (res.status === StatusCodes.OK) {
      this.$router.replace({ path: '/', append: true })
    }
  }
}
</script>
