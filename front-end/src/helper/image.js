export default function getImageBySlug(slug) {
  const baseUrl = 'http://localhost:5173/'
  const defaultImage = +'img/default.png'
  //check if slug is undefined
  if (!slug) {
    return defaultImage ?? baseUrl + 'img/default.jpg'
  }

  let name = slug.replace(/\s/g, '-') ?? ''

  const image = baseUrl + 'img/' + name.toLowerCase() + '.png'

  return image
}
