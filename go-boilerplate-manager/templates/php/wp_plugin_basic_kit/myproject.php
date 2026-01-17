<?php
/**
 * Plugin Name: PROJECT_NAME_PLACEHOLDER
 * Description: PROJECT_DESCRIPTION_PLACEHOLDER
 * Version: 0.1.0
 * Author: AUTHOR_NAME_PLACEHOLDER
 * Text Domain: myproject
 * Domain Path: /languages
 * Requires at least: 6.0
 * Requires PHP: 7.4
 * License: GPL-2.0-or-later
 */

if (!defined('ABSPATH')) {
    exit;
}

require_once plugin_dir_path(__FILE__) . 'includes/class-loader.php';
require_once plugin_dir_path(__FILE__) . 'includes/class-activator.php';
require_once plugin_dir_path(__FILE__) . 'includes/class-deactivator.php';
require_once plugin_dir_path(__FILE__) . 'includes/class-plugin.php';

register_activation_hook(__FILE__, array('MyProject_Activator', 'activate'));
register_deactivation_hook(__FILE__, array('MyProject_Deactivator', 'deactivate'));

function myproject_run() {
    $plugin = new MyProject_Plugin();
    $plugin->run();
}
myproject_run();
