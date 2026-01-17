# PROJECT_NAME_PLACEHOLDER

PROJECT_DESCRIPTION_PLACEHOLDER

## Quick start

1. Copy the folder into `wp-content/plugins/`.
2. Activate the plugin in the WordPress admin.
3. Open Settings -> PROJECT_NAME_PLACEHOLDER to see the example screen.

## Structure

- `includes/` core plugin logic, loader, activation, deactivation
- `admin/` admin-only code, styles, scripts
- `public/` front-end code, styles, scripts

## Defaults

- Text domain: `myproject`
- Shortcode: `[myproject_example]`

## Customize

- Update `admin/class-admin.php` to build settings screens.
- Update `public/class-public.php` to add shortcodes and front-end output.
- Replace placeholder strings in `uninstall.php` if you add stored options.

## Support

- Author: AUTHOR_NAME_PLACEHOLDER
- Email: AUTHOR_EMAIL_PLACEHOLDER
