import { Property } from "../../components/propertyGrid/property";
import { ITranslations } from "../translations/translations-contract";
import BaseReportItemProperties from "./baseReportItemProperties";

export default class TextReportItemProperties extends BaseReportItemProperties {
  constructor(translations?: Partial<ITranslations>) {
    super(translations);
  }
  private _text = "";
  private _binding = "";
  private _format = "";

  get text() {
    return this._text;
  }
  get binding() {
    return this._binding;
  }
  get format() {
    return this._format;
  }

  set text(value: string) {
    const oldValue = this.text;
    this._text = value;
    this.emitOnChange("text", value, oldValue);
  }
  set binding(value: string) {
    const oldValue = this.binding;
    this._binding = value;
    this.emitOnChange("binding", value, oldValue);
  }
  set format(value: string) {
    const oldValue = this.format;
    this._format = value;
    this.emitOnChange("format", value, oldValue);
  }

  getPropertyDefinitions(): Property[] {
    return [
      {
        field: "text",
        label: this.translations?.text ?? "Text",
        type: "string",
      },
      {
        field: "binding",
        label: this.translations?.binding ?? "Binding",
        type: "string",
      },
      {
        field: "format",
        label: this.translations?.format ?? "Format",
        type: "string",
      },
      ...super.getPropertyDefinitions(),
    ];
  }
}
