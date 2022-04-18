export const dateToLapsedTime = (date) => {
  const seconds = Math.floor((new Date() - new Date(date)) / 1000)
  let interval = seconds / 31536000
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' years ago'
  }
  interval = seconds / 2592000
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' months ago'
  }
  interval = seconds / 604800
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' weeks ago'
  }
  interval = seconds / 86400
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' days ago'
  }
  interval = seconds / 3600
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' hours ago'
  }
  interval = seconds / 60
  if (interval > 1) {
    return 'Edited ' + Math.floor(interval) + ' minutes ago'
  }
  return 'Edited recently'
}
