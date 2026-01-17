<?php

if (!defined('ABSPATH')) {
    exit;
}

class MyProject_Public {
    private $plugin_name;
    private $version;

    public function __construct($plugin_name, $version) {
        $this->plugin_name = $plugin_name;
        $this->version = $version;
    }

    public function enqueue_styles() {
        wp_enqueue_style(
            $this->plugin_name . '-public',
            plugin_dir_url(__FILE__) . 'css/public.css',
            array(),
            $this->version,
            'all'
        );
    }

    public function enqueue_scripts() {
        wp_enqueue_script(
            $this->plugin_name . '-public',
            plugin_dir_url(__FILE__) . 'js/public.js',
            array('jquery'),
            $this->version,
            true
        );
    }

    public function register_shortcodes() {
        add_shortcode('myproject_example', array($this, 'render_example'));
    }

    public function render_example($atts = array(), $content = null) {
        $atts = shortcode_atts(
            array(
                'label' => __('Hello from PROJECT_NAME_PLACEHOLDER', 'myproject'),
            ),
            $atts,
            'myproject_example'
        );

        return '<div class="myproject-example">' . esc_html($atts['label']) . '</div>';
    }
}
