<template>
  <card title="あなたの勤怠グループを作成しましょう">
    <form @submit.prevent="submit">
      <div class="mb3">
        <form-label id="group_name" label="Group Name" />
        <input-field id="group_name" v-model="groupName" type="text" />
      </div>
      <div class="mb3">
        <form-label id="login_id" label="Administrator ID" />
        <input-field id="login_id" v-model="loginId" type="text" />
      </div>
      <div class="mb3">
        <form-label id="password" label="Administrator Password" />
        <input-field id="password" v-model="password" type="password" />
      </div>
      <div class="mb3">
        <form-label id="user_name" label="Administrator Name" />
        <input-field id="user_name" v-model="userName" type="text" />
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
  private groupName = ''
  private loginId = ''
  private password = ''
  private userName = ''

  private async submit(): Promise<void> {
    const res = await this.$axios.post('/group/create', {
      name: this.groupName,
      // eslint-disable-next-line camelcase
      login_id: this.loginId,
      password: this.password,
      // eslint-disable-next-line camelcase
      user_name: this.userName,
    })
    if (res.status === StatusCodes.OK) {
      this.$router.replace({ path: '/', append: true })
    }
  }
}
</script>
