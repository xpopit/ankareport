import { defineStore } from "pinia";
import axios from "axios";

export const useReportStore = defineStore("report", {
  state: () => ({
    reports: [],
    currentReport: null,
  }),
  actions: {
    async fetchReports() {
      const { data } = await axios.get("/api/v1/reports");
      this.reports = data || [];
    },
    async fetchReport(id: string) {
      const { data } = await axios.get(`/api/v1/reports/${id}`);
      this.currentReport = data;
      return data;
    },
    async createReport(report: any, definition: any) {
      const { data } = await axios.post("/api/v1/reports", {
        report,
        definition,
      });
      return data;
    },
    async updateReport(id: string, report: any, definition: any) {
      const { data } = await axios.put(`/api/v1/reports/${id}`, {
        report,
        definition,
      });
      return data;
    },
    async deleteReport(id: string) {
      await axios.delete(`/api/v1/reports/${id}`);
    },
    async cloneReport(id: string) {
      const { data } = await axios.post(`/api/v1/reports/${id}/clone`);
      return data;
    },
  },
});
