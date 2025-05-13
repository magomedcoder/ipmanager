<script setup lang="ts">
import { client } from '@/api/client'
import { UserService } from '@/api/pb/user_pb'
import { ConnectError } from '@connectrpc/connect'

const userService = client(UserService)

try {
  const { accessToken } = await userService.login({
    username: '',
    password: ''
  })
  console.log(accessToken)
} catch (err) {
  if (err instanceof ConnectError) {
    console.log(err.message)
  }
}
</script>

<template>

</template>

<style scoped>

</style>
