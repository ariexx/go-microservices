export default {
  getImageBySlug(slug) {
    //get name of product then replace space with dash
    let name = slug.replace(/\s+/g, '-')

    return "img/" + name.toLowerCase() + '.png'
  }
}