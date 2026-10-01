import { QueryEngine } from "./query-engine";
import { Dataset } from "../schema/dataset";

describe("QueryEngine", () => {
  const mockDataset: Dataset = {
    id: "1",
    name: "Sales Data",
    columns: [
      { name: "product", type: "string" },
      { name: "category", type: "string" },
      { name: "amount", type: "number" },
      { name: "status", type: "string" },
    ],
    data: [
      {
        product: "Laptop",
        category: "Electronics",
        amount: 1200,
        status: "Active",
      },
      { product: "Desk", category: "Furniture", amount: 300, status: "Active" },
      {
        product: "Chair",
        category: "Furniture",
        amount: 150,
        status: "Inactive",
      },
      {
        product: "Monitor",
        category: "Electronics",
        amount: 400,
        status: "Active",
      },
    ],
  };

  it("should return all data when no options are provided", () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({});
    expect(result.length).toBe(4);
  });

  it('should filter data using "eq" operator', () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      filters: [{ field: "category", operator: "eq", value: "Furniture" }],
    });
    expect(result.length).toBe(2);
    expect(result.every((r) => r.category === "Furniture")).toBe(true);
  });

  it('should filter data using "gt" operator', () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      filters: [{ field: "amount", operator: "gt", value: 350 }],
    });
    expect(result.length).toBe(2);
    expect(result.find((r) => r.product === "Laptop")).toBeDefined();
    expect(result.find((r) => r.product === "Monitor")).toBeDefined();
  });

  it('should filter data using "contains" operator', () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      filters: [{ field: "product", operator: "contains", value: "top" }],
    });
    expect(result.length).toBe(1);
    expect(result[0].product).toBe("Laptop");
  });

  it("should sort data in ascending order", () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      sort: [{ field: "amount", direction: "asc" }],
    });
    expect(result[0].product).toBe("Chair"); // 150
    expect(result[3].product).toBe("Laptop"); // 1200
  });

  it("should sort data in descending order", () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      sort: [{ field: "amount", direction: "desc" }],
    });
    expect(result[0].product).toBe("Laptop"); // 1200
    expect(result[3].product).toBe("Chair"); // 150
  });

  it("should apply multiple filters", () => {
    const engine = new QueryEngine(mockDataset);
    const result = engine.execute({
      filters: [
        { field: "category", operator: "eq", value: "Electronics" },
        { field: "amount", operator: "lt", value: 1000 },
      ],
    });
    expect(result.length).toBe(1);
    expect(result[0].product).toBe("Monitor");
  });
});
