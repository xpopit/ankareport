import { defineStore } from 'pinia'

// Temporary mock UUID generator
const generateId = () => Math.random().toString(36).substring(2, 9)

export const useReportStore = defineStore('report', {
  state: () => ({
    reports: [] as any[],
    currentReport: null as any,
  }),
  actions: {
    loadFromLocalStorage() {
      const data = localStorage.getItem('xerp_reports')
      if (data) {
        this.reports = JSON.parse(data)
      } else {
        this.reports = []
      }
    },
    saveToLocalStorage() {
      localStorage.setItem('xerp_reports', JSON.stringify(this.reports))
    },
    fetchReports() {
      this.loadFromLocalStorage()
      return this.reports
    },
    fetchReport(id: string) {
      this.loadFromLocalStorage()
      const report = this.reports.find(r => r.id === id)
      this.currentReport = report || null
      return report
    },
    createReport(reportInput: any, definition: any) {
      this.loadFromLocalStorage()
      const newReport = {
        id: generateId(),
        ...reportInput,
        definition: definition,
        currentVersion: 1,
        status: 'DRAFT',
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString()
      }
      this.reports.push(newReport)
      this.saveToLocalStorage()
      return newReport
    },
    updateReport(id: string, reportUpdate: any, definition: any) {
      this.loadFromLocalStorage()
      const index = this.reports.findIndex(r => r.id === id)
      if (index !== -1) {
        this.reports[index] = {
          ...this.reports[index],
          ...reportUpdate,
          definition: definition,
          currentVersion: (this.reports[index].currentVersion || 1) + 1,
          updatedAt: new Date().toISOString()
        }
        this.saveToLocalStorage()
        return this.reports[index]
      }
      throw new Error('Report not found')
    },
    deleteReport(id: string) {
      this.loadFromLocalStorage()
      this.reports = this.reports.filter(r => r.id !== id)
      this.saveToLocalStorage()
    },
    cloneReport(id: string) {
      this.loadFromLocalStorage()
      const original = this.reports.find(r => r.id === id)
      if (!original) throw new Error('Report not found')

      const clone = JSON.parse(JSON.stringify(original))
      clone.id = generateId()
      clone.name = clone.name + ' (Clone)'
      clone.currentVersion = 1
      clone.status = 'DRAFT'
      clone.createdAt = new Date().toISOString()
      clone.updatedAt = new Date().toISOString()

      this.reports.push(clone)
      this.saveToLocalStorage()
      return clone
    }
  }
})
