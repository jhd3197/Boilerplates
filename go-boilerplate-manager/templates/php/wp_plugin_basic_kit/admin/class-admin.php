<?php

if (!defined('ABSPATH')) {
    exit;
}

class MyProject_Admin {
    private $plugin_name;
    private $version;

    public function __construct($plugin_name, $version) {
        $this->plugin_name = $plugin_name;
        $this->version = $version;
    }

    public function add_settings_page() {
        add_options_page(
            'PROJECT_NAME_PLACEHOLDER',
            'PROJECT_NAME_PLACEHOLDER',
            'manage_options',
            $this->plugin_name,
            array($this, 'render_settings_page')
        );
    }

    public function render_settings_page() {
        if (!current_user_can('manage_options')) {
            return;
        }
        ?>
        <div class="wrap">
            <h1><?php echo esc_html(get_admin_page_title()); ?></h1>
            <p><?php echo esc_html__('Settings for PROJECT_NAME_PLACEHOLDER.', 'myproject'); ?></p>
        </div>
        <?php
    }

    public function enqueue_styles() {
        wp_enqueue_style(
            $this->plugin_name . '-admin',
            plugin_dir_url(__FILE__) . 'css/admin.css',
            array(),
            $this->version,
            'all'
        );
    }

    public function enqueue_scripts() {
        wp_enqueue_script(
            $this->plugin_name . '-admin',
            plugin_dir_url(__FILE__) . 'js/admin.js',
            array('jquery'),
            $this->version,
            false
        );
    }
}
