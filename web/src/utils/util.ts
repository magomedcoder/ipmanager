export const getLastOctet = (ipAddress) => {
  const parts = ipAddress.split('.')
  return '.' + parts[3]
}
