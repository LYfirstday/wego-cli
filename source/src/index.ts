#!/usr/bin/env node

import {
  initYamlFile,
  requestRemoteTemplates,
  TemplateType,
} from "@lygo/wego-rs";
import { Command } from "commander";
const packageJson = require("./../package.json");

const program = new Command();

program
  .name("Weee frontend templates utils")
  .description(
    "CLI to download the frontend pages、components、projects template"
  )
  .version(packageJson.version);

program
  .command("init")
  .description("Generate a wego.yaml config file")
  .action(() => {
    initYamlFile();
  });

program
  .command("show")
  .description("Display the pages、components、projects templates")
  .option("-j --projects", "Display the projects templates")
  .option("-p --pages", "Display the pages templates")
  .option("-c --components", "Display the components templates")
  .action((select_obj) => {
    if (select_obj.projects) {
      requestRemoteTemplates(TemplateType.Project);
    } else if (select_obj.pages) {
      requestRemoteTemplates(TemplateType.Pages);
    } else if (select_obj.components) {
      requestRemoteTemplates(TemplateType.Components);
    }
  });

program.parse();
