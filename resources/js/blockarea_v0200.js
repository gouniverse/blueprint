/*
 * blockarea_v0200.js
 *
 * This file is copyrighted and not released to the public domain.
 * No part of this file may be redistributed in any form without
 * prior written permission by the author.
 *
 * Copyright 2024 Milan Lesichkov
 *
 * All rights reserved.
 */

var BlockAreaCode = {
  type: "Code",
  name: "Code",
  render: function (block) {
    if (!block.Attributes.Code) {
      return "Not set up yet";
    }
    var html = "";
    html += "<p>Language: " + block.Attributes.Language + "</p>";
    html += "<pre><code>" + block.Attributes.Code + "</code></pre>";
    return html;
  },
  template: function () {
    var html = "";
    html += '<div class="form-group">';
    html += "<label>Language</label>";
    html += '<select name="Language" class="form-control">';
    html += "<option></option>";
    html += '<option value="bash">Bash</option>';
    html += '<option value="css">CSS</option>';
    html += '<option value="c">C</option>';
    html += '<option value="cpp">C++</option>';
    html += '<option value="csharp">C#</option>';
    html += '<option value="golang">Golang</option>';
    html += '<option value="html">HTML</option>';
    html += '<option value="javascript">JavaScript</option>';
    html += '<option value="java">Java</option>';
    html += '<option value="php">PHP</option>';
    html += '<option value="python">Python</option>';
    html += '<option value="sql">SQL</option>';
    html += '<option value="xhtml">XHTML</option>';
    html += '<option value="xml">XML</option>';
    html += '<option value="other">Other</option>';
    html += "</select>";
    html += "</div>";
    html += '<div class="form-group">';
    html += "<label>Code</label>";
    html +=
      '<textarea class="form-control" name="Code" placeholder="Please enter code here"></textarea>';
    html += "</div>";
    return html;
  },
};

var BlockAreaRawHtml = {
  type: "RawHtml",
  name: "Raw HTML",
  render: function (block) {
    if (!block.Attributes.Text) {
      return "Not set up yet";
    }
    var html = "";
    html += "<p>Raw HTML code</p>";
    return html;
  },
  template: function () {
    var html = "";
    html +=
      '<textarea class="form-control" name="Text" placeholder="Your text here" value=""></textarea>';
    return html;
  },
};

var BlockAreaHeading = {
  type: "Heading",
  name: "Heading",
  render: function (block) {
    if (!block.Attributes.Text) {
      return "Not set up yet";
    }
    var html = "";
    html +=
      "<h" +
      block.Attributes.Level +
      ">" +
      block.Attributes.Text +
      "</h" +
      block.Attributes.Level +
      ">";
    return html;
  },
  template: function () {
    var html = "";
    html += '<div class="form-group">';
    html += "<label>Level</label>";
    html += '<select name="Level" class="form-control">';
    html += "<option></option>";
    html += '<option value="1">Heading 1</option>';
    html += '<option value="2">Heading 2</option>';
    html += '<option value="3">Heading 3</option>';
    html += '<option value="4">Heading 4</option>';
    html += '<option value="5">Heading 5</option>';
    html += '<option value="6">Heading 6</option>';
    html += "</select>";
    html += "</div>";
    html += '<div class="form-group">';
    html += "<label>Text</label>";
    html +=
      '<input class="form-control" name="Text" value="" placeholder="Please enter heading text here" />';
    html += "</div>";
    return html;
  },
};
var BlockAreaImage = {
  type: "Image",
  name: "Image",
  render: function (block) {
    if (!block.Attributes.Url) {
      return "Not set up yet";
    }
    var html = "";
    html +=
      '<img src="' + block.Attributes.Url + '" style="max-height:100px;" />';
    return html;
  },
  template: function () {
    var html = "";
    html += '<div class="form-group">';
    html += '<div class="form-group">';
    html += "<label>Url</label>";
    html +=
      '<textarea class="form-control" name="Url" placeholder="Please enter URL for image here"></textarea>';
    html += "</div>";
    return html;
  },
};
var BlockAreaText = {
  type: "Text",
  name: "Text",
  render: function (block) {
    if (!block.Attributes.Text) {
      return "Not set up yet";
    }
    var html = "";
    html += "<p>" + block.Attributes.Text + "</p>";
    return html;
  },
  template: function () {
    var html = "";
    html +=
      '<textarea class="form-control" name="Text" placeholder="Your text here" value=""></textarea>';
    return html;
  },
};

function BlockArea(blockAreaId) {
  var self = this;
  var blockAreaId = blockAreaId;
  var blocks = [];
  var blockDefinitions = [];
  var editor = $("#" + blockAreaId);
  var isTextArea = false;
  var textAreaId = null;

  //var urlBlock = '';
  var blockParentId = "";
  //var urlBlockTemplate = '';
  //var urlBlockCreate = '';
  //var urlBlockDelete = '';
  //var urlBlockList = '';
  //var urlBlockSave = '';
  //var urlBlockListSequenceSave = '';
  //var urlBlockDelete = '';

  function _block(block) {
    var blockDefinition = "BlockArea" + block.Type;
    var blockRendered = window[blockDefinition].render(block);
    var html = "";
    html +=
      '<div class="block block-id-' +
      block.Id +
      '" data-id="' +
      block.Id +
      '">';
    html += '  <div class="block-header">';
    html += '    <span class="handle" title="Move">&varr;</span>';
    html +=
      '    <span class="title">' + block.Type + " - " + block.Id + "</span>";
    html +=
      '    <span class="setting" title="Settings" onclick="$.publish(\'' +
      blockAreaId +
      "_Attributes',{Id:'" +
      block.Id +
      '\'})" style="float:right;text-decoration:underline;">';
    html += "      Settings";
    html += "    </span>";
    html +=
      '    <span class="delete" title="Delete block" onclick="$.publish(\'' +
      blockAreaId +
      "_Delete',{Id:'" +
      block.Id +
      '\'})" style="float:right;text-decoration:underline;">';
    html += "      Delete";
    html += "    </span>";
    html += " </div>";
    html += ' <div class="block-settings">';
    html += " </div>";
    html += ' <div class="block-body">';
    html += blockRendered;
    html += " </div>";
    html += ' <div class="block-footer">';
    html += " </div>";
    html += "</div>";
    return html;
  }

  /**
   * Find a block index by the block ID
   * @returns {void}
   */
  function _findBlockIndexById(blockId) {
    for (var i = 0; i < blocks.length; i++) {
      if (blocks[i].Id === blockId) {
        return i;
      }
    }
    return null;
  }

  /**
   * Find a block index by the block ID
   * @returns {void}
   */
  function _findBlockById(blockId) {
    var index = _findBlockIndexById(blockId);
    if (index === null) {
      return null;
    }
    if (blocks[index]) {
      return blocks[index];
    }
    return null;
  }

  function _formPopulate(form, data) {
    // DEBUG: console.log(data);
    $.each(data, function (key, value) {
      // key = 'Attribute[' + key + ']';
      // console.log(key);
      var name_key = key.split("[").join("\\[").split("]").join("\\]");
      var $ctrl = $("[name=" + name_key + "]", form);
      if (!$ctrl) {
        return; // Not found
      }
      var type = $ctrl.attr("type");
      var tag = $ctrl.prop("tagName").toLowerCase();
      //console.log(name_key + ":" + value);
      //console.log(tag);

      if (tag === "input") {
        var type = $ctrl.attr("type");
        if (type === "radio" || type === "checkbox") {
          $ctrl.each(function () {
            if ($(this).attr("value") === value) {
              $(this).attr("checked", value);
            }
          });
        } else {
          $ctrl.val(value);
        }
      }
      if (tag === "select") {
        $ctrl.val(value);
      }
      if (tag === "textarea") {
        $ctrl.val(value);
      }
    });
  }

  /**
   * Generates a human friendly unique identifier
   * @returns {void}
   */
  function _huid(options) {
    var options = typeof options === "undefined" ? {} : options;
    var useDashes =
      typeof options.useDashes === "undefined" ? true : options.useDashes;
    var addMilliseconds =
      typeof options.addMilliseconds === "undefined"
        ? true
        : options.addMilliseconds;

    var date = new Date();
    var year = date.getUTCFullYear();
    var month = 1 + date.getUTCMonth();
    month = ("00" + month).slice(-2);
    var day = date.getUTCDate();
    day = ("00" + day).slice(-2);
    var hours = date.getUTCHours();
    hours = ("00" + hours).slice(-2);
    var minutes = date.getUTCMinutes();
    minutes = ("00" + minutes).slice(-2);
    var seconds = date.getUTCSeconds();
    seconds = ("00" + seconds).slice(-2);
    var milliseconds = date.getUTCMilliseconds();
    milliseconds = ("0000" + milliseconds).slice(-4);
    var uuid = year;
    uuid += "" + month;
    uuid += "" + day;
    if (useDashes === true) {
      uuid += "-";
    }
    uuid += "" + hours;
    uuid += "" + minutes;
    uuid += "" + seconds;
    if (useDashes === true && addMilliseconds === true) {
      uuid += "-";
    }
    if (addMilliseconds === true) {
      uuid += "" + milliseconds;
    }
    return uuid;
  }

  /**
   * Called when the blocks are reordered
   * @returns {void}
   */
  function _onReordered() {
    $.publish(blockAreaId + "_TextAreaUpdate");

    var blockSequenceMap = {};

    editor.find(".blocks .block").each(function (index) {
      var id = $(this).data("id");
      blockSequenceMap[id] = index + 1;
    });

    var newBlocks = [];
    $.each(blockSequenceMap, function (blockId, sequence) {
      var block = _findBlockById(blockId);
      if (block === null) {
        return;
      }
      block.Sequence = sequence;
      newBlocks.push(block);
    });

    blocks = newBlocks;

    $.publish(blockAreaId + "_TextAreaUpdate");
  }

  /**
   * Called when the textarea content must be updated
   * @returns {void}
   */
  function _onTextAreaUpdate() {
    if (isTextArea == false) {
      return;
    }
    console.log("onTextAreaUpdate");
    $("#" + textAreaId).val(JSON.stringify(blocks));
    $("#" + textAreaId).trigger("change");
    $("#" + textAreaId).trigger("input");

    // Update via native event (required for VueJS)
    const event = new Event("input", { bubbles: true });
    document.getElementById(textAreaId).dispatchEvent(event);
  }

  /**
   * Generates the interface for the toolbar
   * @returns {void}
   */
  function _toolbar() {
    var html = "";
    html += '<div class="toolbar">';
    for (var i = 0; i < blockDefinitions.length; i++) {
      var blockDefinition = blockDefinitions[i];
      var type = blockDefinition.type;
      var name = blockDefinition.name;
      html +=
        '  <button type="button" onclick="$.publish(\'' +
        blockAreaId +
        "_Create',{type:'" +
        type +
        '\'});"><i class="fa fa-plus-circle"></i> ' +
        name +
        "</button>";
    }
    html += "</div>";
    return html;
  }

  /**
   * Creates a new block
   * @returns {void}
   */
  this.blockCreate = function () {
    var block = arguments[1];
    var blockType = block.type;
    var blockSequence = blocks.length + 1;
    var block = {
      Id: _huid({ useDashes: false }),
      ParentId: blockParentId,
      Type: blockType,
      Sequence: blockSequence,
      Attributes: {},
    };
    blocks.push(block);
    $.publish(blockAreaId + "_LoadBlocks");
  };

  /**
   * Deletes a block by ID
   * @returns {void}
   */
  this.blockDelete = function () {
    var block = arguments[1];
    var blockId = block.Id;
    var newBlocks = [];
    for (var i = 0; i < blocks.length; i++) {
      if (blocks[i].Id === blockId) {
        continue;
      }
      newBlocks.push(blocks[i]);
    }
    blocks = newBlocks;
    $.publish(blockAreaId + "_LoadBlocks");
    $.publish(blockAreaId + "_TextAreaUpdate");
  };

  this.blockAttributes = function () {
    var arguments = arguments[1];
    var blockId = arguments.Id;
    var block = blocks[_findBlockIndexById(blockId)];
    var blockType = block.Type;
    var blockAttributes = block.Attributes;
    var blockDiv = editor.find(".block-id-" + blockId);
    var settingsDiv = blockDiv.find(".block-settings");
    var blockDefinition = "BlockArea" + blockType;
    var template = window[blockDefinition].template();
    var html = "";
    html += template;
    html +=
      '<div style="border-top:1px solid brown;font-size:1px;margin:2px 0px;">&nbsp;</div>';
    html +=
      '<button type="button" class="btn btn-success" style="margin:0px 10px 0px 0px;" onclick="$.publish(\'' +
      blockAreaId +
      "_AttributesSave',{Id:'" +
      block.Id +
      "'})\">OK</button>";
    html +=
      '<button type="button" class="btn btn-default" onclick="$(this).parent().html(\'\')">Close</button>';
    settingsDiv.html(html);
    var form = $(".block-id-" + blockId + " .block-settings").eq(0);
    _formPopulate(form, blockAttributes);
  };

  this.blockAttributesSave = function () {
    var block = arguments[1];
    var blockId = block.Id;
    var blockIndex = _findBlockIndexById(blockId);
    var blockDiv = editor.find(".block-id-" + blockId);
    var settingsDiv = blockDiv.find(".block-settings");
    var attributes = $(
      ".block-id-" + blockId + " .block-settings :input"
    ).serializeObject();
    blocks[blockIndex].Attributes = attributes;
    settingsDiv.html("Saved");

    $.publish(blockAreaId + "_LoadBlocks");
    $.publish(blockAreaId + "_TextAreaUpdate");

    setTimeout(function () {
      settingsDiv.html("");
    }, 3000);
  };

  /**
   * Sets the blocks for the block area
   * @returns {undefined}
   */
  this.setBlocks = function (blockList) {
    blocks = blockList;
  };

  /**
   * Sets the blocks for the block area
   * @returns {undefined}
   */
  this.getBlocks = function () {
    return blocks;
  };

  /**
   * Gets the blocks
   * @returns {undefined}
   */
  this.getBlocksAsJson = function () {
    return JSON.stringify(blocks);
  };

  function displayBlocks() {
    editor.html("");
    var toolbar = _toolbar();
    var html = "";
    html += toolbar;
    html += '<div class="blocks">';
    for (var i = 0; i < blocks.length; i++) {
      var block = blocks[i];
      html += _block(block);
    }
    html += "</div>";
    editor.append(html);
  }

  this.loadBlocks = function () {
    displayBlocks();

    editor.find(".blocks").sortable({
      placeholder: "ui-state-highlight",
      handle: ".handle",
      cursor: "pointer",
      stop: function (event, ui) {
        $.publish(blockAreaId + "_Reordered");
      },
    });
    skinApply();
  };

  this.registerBlock = function (blockDefinition) {
    blockDefinitions.push(blockDefinition);
  };

  this.setParentId = function (parentId) {
    blockParentId = parentId;
  };

  //    this.setUrlBlockDelete = function (url) {
  //        urlBlockDelete = url;
  //    };
  //
  //    this.setUrlBlockSave = function (url) {
  //        urlBlockSave = url;
  //    };
  //    this.setUrlBlockCreate = function (url) {
  //        urlBlockCreate = url;
  //    };
  //    this.setUrlBlockList = function (url) {
  //        urlBlockList = url;
  //    };
  //    this.setUrlBlockListSequenceSave = function (url) {
  //        urlBlockListSequenceSave = url;
  //    };
  //    this.setUrlBlockTemplate = function (url) {
  //        urlBlockTemplate = url;
  //    };
  function skinApply() {
    editor.css({
      border: "1px solid #666",
      //'min-height': '300px',
      "font-family": "verdana",
      "font-size": "12px",
    });

    editor.find(".toolbar").css({
      border: "1px solid darkslategray",
      background: "cornflowerblue",
      padding: "5px",
    });

    editor.find(".blocks").css({
      background: "azure",
      padding: "5px",
    });

    editor.find(".block").css({
      border: "1px solid lightcoral",
      position: "relative",
      padding: "0px",
      margin: "5px 0px 10px 0px",
    });

    editor.find(".block .block-header").css({
      background: "cadetblue",
      color: "white",
      padding: "5px",
    });

    editor.find(".block .block-settings").css({
      background: "lightsteelblue",
      padding: "5px",
    });

    editor.find(".block .handle").css({
      display: "inline-block",
      width: "20px",
      height: "20px",
      color: "orangered",
      background: "white",
      "border-radius": "10px",
      "text-align": "center",
      margin: "0px 5px",
      "font-size": "12px",
      "line-height": "19px",
      padding: "0px",
      cursor: "pointer",
      "font-weight": "800",
    });
    editor.find(".block .title").css({
      "font-weight": "bold",
    });

    editor.find(".block .setting").css({
      cursor: "pointer",
      float: "right",
    });

    editor.find(".block .delete").css({
      cursor: "pointer",
      float: "right",
      margin: "0px 5px 0px 0px",
    });
  }

  /**
   * Initializes the block area
   * @returns {undefined}
   */
  this.init = function () {
    /* Find if called for textarea */
    var tag = $("#" + blockAreaId)
      .prop("tagName")
      .toLowerCase();
    if (tag === "textarea") {
      isTextArea = true;
    }

    /* if textarea create the blockarea, and hide the textarea */
    if (isTextArea) {
      textAreaId = blockAreaId;
      blockAreaId = textAreaId + "_BlockArea";
      $("#" + textAreaId)
        .parent()
        .prepend('<div id="' + blockAreaId + '">&nbsp;</div>');
      editor = $("#" + blockAreaId); // Replace editor
      var textareaValue = $.trim($("#" + textAreaId).val());
      if (blocks.length < 1) {
        if (textareaValue != "") {
          if (
            textareaValue.substr(0, 1) == "[" &&
            textareaValue.substr(-1) == "]"
          ) {
            blocks = JSON.parse(textareaValue);
          } else {
            console.log("Textarea value not JSON");
            return false;
          }
        }
      }
      $("#" + textAreaId).hide();
    }

    /* Display the blocks */
    this.loadBlocks();

    /* Attach event listeners */
    $.subscribe(blockAreaId + "_Create", this.blockCreate);
    $.subscribe(blockAreaId + "_Delete", this.blockDelete);
    $.subscribe(blockAreaId + "_Attributes", this.blockAttributes);
    $.subscribe(blockAreaId + "_AttributesSave", this.blockAttributesSave);
    $.subscribe(blockAreaId + "_LoadBlocks", this.loadBlocks);
    $.subscribe(blockAreaId + "_Reordered", _onReordered);
    $.subscribe(blockAreaId + "_TextAreaUpdate", _onTextAreaUpdate);

    /* Update textarea */
    $.publish(blockAreaId + "_TextAreaUpdate");
  };
}

(function ($) {
  var o = $({});
  $.subscribe = function () {
    o.on.apply(o, arguments);
  };
  $.unsubscribe = function () {
    o.off.apply(o, arguments);
  };
  $.publish = function () {
    o.trigger.apply(o, arguments);
  };
  $.fn.serializeObject = function () {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function () {
      if (o[this.name] !== undefined) {
        if (!o[this.name].push) {
          o[this.name] = [o[this.name]];
        }
        o[this.name].push(this.value || "");
      } else {
        o[this.name] = this.value || "";
      }
    });
    return o;
  };
})(jQuery);
