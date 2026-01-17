<?php

if (!defined('ABSPATH')) {
    exit;
}

class MyProject_Plugin {
    protected $loader;
    protected $plugin_name;
    protected $version;

    public function __construct() {
        $this->plugin_name = 'myproject';
        $this->version = '0.1.0';

        $this->load_dependencies();
        $this->define_admin_hooks();
        $this->define_public_hooks();
    }

    private function load_dependencies() {
        require_once plugin_dir_path(dirname(__FILE__)) . 'admin/class-admin.php';
        require_once plugin_dir_path(dirname(__FILE__)) . 'public/class-public.php';

        $this->loader = new MyProject_Loader();
    }

    private function define_admin_hooks() {
        $plugin_admin = new MyProject_Admin($this->plugin_name, $this->version);

        $this->loader->add_action('admin_menu', $plugin_admin, 'add_settings_page');
        $this->loader->add_action('admin_enqueue_scripts', $plugin_admin, 'enqueue_styles');
        $this->loader->add_action('admin_enqueue_scripts', $plugin_admin, 'enqueue_scripts');
    }

    private function define_public_hooks() {
        $plugin_public = new MyProject_Public($this->plugin_name, $this->version);

        $this->loader->add_action('wp_enqueue_scripts', $plugin_public, 'enqueue_styles');
        $this->loader->add_action('wp_enqueue_scripts', $plugin_public, 'enqueue_scripts');
        $this->loader->add_action('init', $plugin_public, 'register_shortcodes');
    }

    public function run() {
        $this->loader->run();
    }

    public function get_plugin_name() {
        return $this->plugin_name;
    }

    public function get_version() {
        return $this->version;
    }
}
