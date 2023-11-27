const { readdirSync } = require('fs');

module.exports = {
  title: 'str',
  description: 'A fluent string manipulation library in Go',
  head: [
    ['link', {
      rel: 'icon',
      type: 'image/png',
      href: 'favicon.png',
    }],
  ],
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'GitHub', link: 'https://github.com/ecrmnn/str' },
    ],
    sidebar: [{
      title: 'API',
      collapsable: false,
      children: readdirSync('api', 'utf-8').map(file => [
        `/api/${file}`,
        file
          .replace('.md', '')
          .split('_')
          .map(word => word.charAt(0).toUpperCase() + word.slice(1))
          .join(''),
      ]),
    }],
  },
};
