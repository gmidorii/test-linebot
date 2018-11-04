
export const sendMessage = async (message) => {
  try {
      await liff.sendMessages([
        {
          type: 'text',
          text: message
        }
      ])

      liff.closeWindow()
  } catch (error) {
    window.alert(error)
  }
}