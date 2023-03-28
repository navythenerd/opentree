# OpenTree
OpenTree is an easy to configure, deployable and self-hosted LinkTree alternative written in go.

## Configuration
OpenTree uses an easy to compose JSON config file. Have a look at the provided `config.json`.

### Links
The link array is used to compose the links, with a link entry for each linked website. Links can be enabled/disabled through the `enabled` field. 
```
"links": [
    {
        "enabled": true,
        "title": "Foo",
        "url": "https://foo.world"
    },
    {
        "enabled": true,
        "title": "Bar",
        "url": "https://bar.world"
    },
    {
        "enabled": false,
        "title": "Disabled link",
        "url": "https://disabled-link.world"
    },
]
```

You mave have noticed another link field `website`. This link field is intended to be used as your main website link, which is displayed in the header, right beneath the avatar.

### Avatar
By setting the `avatar` field to `true` the template includes a picture in the header. The avatar has to be saved as `img/avatar.jpg` relative to the given `assets` directory.

## Themes
The colour scheme used for styling is located under `assets/scss/themes/_default.scss`. The scheme can be customized by either changing the file directly or creating a new theme file and including it in the `asstes/scss/app.scss`

```
@use "themes/default" as theme;
```

## Template
The go backend uses the `html/template` package to parse and execute the given templates in the `views` directory. For layout changes a good starting point is the master template given with `views/master.gohtml`.

## CSS Framework & Further Theming
Special thanks to the Bulma devs for providing the base css framework (https://github.com/jgthms/bulma).
Have a look at their documentation for full customization.

## Deployment
You can use the provided Dockerfile to deploy the application. If you have changed the default port `8000`, make sure to expose the right port.
