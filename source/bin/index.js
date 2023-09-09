#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const wego_rs_1 = require("@lygo/wego-rs");
const commander_1 = require("commander");
const program = new commander_1.Command();
program
    .name("Weee frontend templates utils")
    .description("CLI to download the frontend pages、components、projects template")
    .version("0.0.1");
program
    .command("init")
    .description("Generate a wego.yaml config file")
    .action(() => {
    (0, wego_rs_1.initYamlFile)();
});
program
    .command("show")
    .description("Display the pages、components、projects templates")
    .option("-j --projects", "Display the projects templates")
    .option("-p --pages", "Display the pages templates")
    .option("-c --components", "Display the components templates")
    .action((select_obj) => {
    if (select_obj.projects) {
        (0, wego_rs_1.requestRemoteTemplates)(2 /* TemplateType.Project */);
    }
    else if (select_obj.pages) {
        (0, wego_rs_1.requestRemoteTemplates)(0 /* TemplateType.Pages */);
    }
    else if (select_obj.components) {
        (0, wego_rs_1.requestRemoteTemplates)(1 /* TemplateType.Components */);
    }
});
program.parse();
