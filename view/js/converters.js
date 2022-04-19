export const dateToLapsedTime = (_date) => {
  const seconds = Math.floor((new Date() - new Date(_date)) / 1000)
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

export const mimeTypeToAppType = (_mimetype) => {
  if (/image\/*/.test(_mimetype)) {
    return 'image'
  } else if (/video\/*/.test(_mimetype)) {
    return 'video'
  } else if (/audio\/*/.test(_mimetype)) {
    return 'audio'
  } else if (/application\/vnd.ms-excel*/.test(_mimetype) || _mimetype === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' || _mimetype === 'application/vnd.openxmlformats-officedocument.spreadsheetml.template') {
    return 'binder'
  } else if (/application\/vnd\.openxmlformats-officedocument.presentationml\.*/.test(_mimetype) || _mimetype === 'application/vnd.ms-powerpoint') {
    return 'presentation'
  } else if (/application\/vnd\.ms-word\.*/.test(_mimetype) || /application\/vnd\.openxmlformats-officedocument\.wordprocessingml\.*/.test(_mimetype || /application\/msword*/.test(_mimetype))) {
    return 'document'
  } else {
    return 'unknow'
  }
}
