import { Property } from "../../components/propertyGrid/property";
import StyleProperties from "../../core/styleProperties";
import { ITranslations } from "../../core/translations/translations-contract";

const DEFAULT_REPORT_WIDTH = 400;
const MIN_REPORT_WIDTH = 100;

export default class ReportProperties extends StyleProperties {
  constructor(translations?: Partial<ITranslations>) {
    super(undefined, translations);
  }
  private _width = DEFAULT_REPORT_WIDTH;

  get width() {
    return this._width;
  }

  set width(value: number) {
    const oldValue = this.width;
    this._width = Math.max(value, MIN_REPORT_WIDTH);
    this.emitOnChange("width", value, oldValue);
  }

  getPropertyDefinitions(): Property[] {
    return [
      {
        field: "width",
        label: this.translations?.width ?? "Width",
        type: "number",
      },
      ...super.getPropertyDefinitions(),
    ];
  }
}
