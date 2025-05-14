import { useCookies } from '@vueuse/integrations/useCookies'

const cookies = useCookies(['locale'])

const accessToken = 'access_token'

export const isLoggedIn = () => {
  return getAccessToken() != ''
}

export const getAccessToken = () => {
  return cookies.get(accessToken) || ''
}

export const setAccessToken = (token = '', timestamp: number = 2678400): void => {
  const date = new Date(timestamp * 1000)

  cookies.set(accessToken, token, {
    expires: date
  })
}

export const removeAccessToken = () => {
  cookies.remove(accessToken)
}
