import "jquery-ujs";
import "bootstrap/js/dist/util.js";
import "bootstrap/js/dist/collapse.js";
import "bootstrap/js/dist/alert.js";
import "bootstrap/js/dist/dropdown.js";
import "bootstrap/js/dist/tab.js";

import { initEditor, initMiniEditor } from "./admin/utils/editor.js";

$(function() {
    if ($("#editor").length) {
        initEditor("#editor");
    }

    if ($("#mini-editor").length) {
        initMiniEditor("#mini-editor");
    }
});
