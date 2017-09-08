import froalaEditor from "froala-editor/js/froala_editor.min.js";

import "froala-editor/js/plugins/align.min.js";
import "froala-editor/js/plugins/code_beautifier.min.js";
import "froala-editor/js/plugins/code_view.min.js";
import "froala-editor/js/plugins/colors.min.js";
import "froala-editor/js/plugins/emoticons.min.js";
import "froala-editor/js/plugins/entities.min.js";
import "froala-editor/js/plugins/table.min.js";
import "froala-editor/js/plugins/lists.min.js";
import "froala-editor/js/plugins/font_size.min.js";
import "froala-editor/js/plugins/url.min.js";
import "froala-editor/js/plugins/link.min.js";
import "froala-editor/js/plugins/quote.min.js";
import "froala-editor/js/plugins/image.min.js";
import "froala-editor/js/plugins/image_manager.min.js";
import "froala-editor/js/plugins/paragraph_style.min.js";
import "froala-editor/js/plugins/paragraph_format.min.js";
import "froala-editor/js/plugins/line_breaker.min.js";

import "froala-editor/css/froala_editor.css";
import "froala-editor/css/froala_style.css";
import "froala-editor/css/plugins/code_view.min.css";
import "froala-editor/css/plugins/emoticons.min.css";
import "froala-editor/css/plugins/draggable.min.css";
import "froala-editor/css/plugins/image.min.css";
import "froala-editor/css/plugins/image_manager.min.css";
import "froala-editor/css/plugins/line_breaker.min.css";
import "froala-editor/css/plugins/table.min.css";

$.FroalaEditor.DefineIcon("readMore", {
    NAME: "arrows-h"
});

$.FroalaEditor.RegisterCommand("readMore", {
    title: "Add a Read More Break",
    focus: true,
    undo: true,
    refreshAfterCallback: true,
    callback: function() {
        this.html.insert(
            '<div class="froala-read-more-break"><!-- truncate --></div>'
        );
        this.html.cleanEmptyTags();
    }
});

// Initialize
export function initEditor(selector) {
    var csrf = $("meta[name=csrf]").attr("content");

    $(selector)
        .froalaEditor({
            imageUploadURL: "/admin/editor-image-upload",
            imageUploadParams: { _csrf_token: csrf },
            heightMin: 500,
            heightMax: 800,
            toolbarButtons: [
                "bold",
                "italic",
                "underline",
                "strikeThrough",
                "subscript",
                "superscript",
                "fontFamily",
                "fontSize",
                "|",
                "color",
                "emoticons",
                "inlineStyle",
                "paragraphStyle",
                "|",
                "paragraphFormat",
                "align",
                "formatOL",
                "formatUL",
                "outdent",
                "indent",
                "quote",
                "insertHR",
                "insertLink",
                "insertImage",
                "insertVideo",
                "insertFile",
                "insertTable",
                "undo",
                "redo",
                "clearFormatting",
                "selectAll",
                "html",
                "readMore"
            ],
            codeMirrorOptions: {
                indentWithTabs: true,
                lineNumbers: true,
                lineWrapping: true,
                mode: "text/html",
                tabMode: "indent",
                tabSize: 2
            },
            htmlRemoveTags: ["style", "base"]
        })
        .on("froalaEditor.image.removed", function(e, editor, $img) {
            $.ajax({
                method: "DELETE",
                url: "/admin/editor-image-upload",
                data: {
                    src: $img.attr("src"),
                    _csrf_token: csrf
                }
            })
                .done(function(data) {
                    console.log("Deleted Image.");
                })
                .fail(function() {
                    console.error("ERROR! Could not delete image");
                });
        });
}

export function initMiniEditor(selector) {
    $(selector).froalaEditor({
        height: 250,
        charCounterCount: false,
        toolbarButtons: [
            "bold",
            "italic",
            "underline",
            "strikeThrough",
            "color",
            "emoticons",
            "paragraphFormat",
            "align",
            "formatOL",
            "formatUL",
            "indent",
            "outdent",
            "insertLink",
            "undo",
            "redo"
        ]
    });
}
