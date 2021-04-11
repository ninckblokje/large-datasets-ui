<template>
  <div class="row">
    <div class="col-sm-1">
      <form>
        <div class="form-group">
          <input id="sizeInput" class="form-control" type="number" min="1" max="10000000" placeholder="Size"
                 v-model.number="size">
        </div>
        <button class="btn btn-primary" @click.prevent="getDataSet">Get dataset</button>
        <button class="btn btn-danger" @click.prevent="reset">Reset</button>
      </form>
      <label>Time: {{ executionTime }}</label>
    </div>
    <div class="col-sm-11">
      <ag-grid-vue style="width: 55rem; height: 50rem;"
                   class="ag-theme-alpine"
                   :columnDefs="columns"
                   :rowData="pageData">
      </ag-grid-vue>
      <div class="btn-group">
        <button class="btn btn-outline-primary" :disabled="disablePreviousButton" @click="previousPage">Previous
        </button>
        <button class="btn btn-outline-primary" :disabled="disableNextButton" @click="nextPage">Next</button>
      </div>
    </div>
  </div>
</template>

<script>
import {AgGridVue} from "ag-grid-vue"

export default {
  name: "Dataset",
  components: {
    AgGridVue
  },
  data() {
    return {
      columns: [
        {field: 'dateTime'},
        {field: 'value1'},
        {field: 'value2'},
        {field: 'value3'}
      ],
      dataSize: 0,
      endDate: null,
      pageData: [],
      pageIndex: 0,
      pageSize: 50,
      size: 1,
      startDate: null
    }
  },
  computed: {
    disableNextButton() {
      return (this.pageIndex * this.pageSize) + this.pageSize >= this.dataSize
    },
    disablePreviousButton() {
      return this.pageIndex == 0
    },
    executionTime() {
      if (this.startDate == null || this.endDate == null) {
        return 0
      } else {
        return this.endDate.getTime() - this.startDate.getTime()
      }
    }
  },
  methods: {
    getDataSet() {
      this.reset()
      this.startDate = new Date()

      // eslint-disable-next-line no-undef
      GetDataSet(this.size)
          .then(dataSize => {
            this.dataSize = dataSize
            this.populatePage(0)
            this.endDate = new Date()
          })
          .catch(err => console.log(err))
    },
    nextPage() {
      this.populatePage(this.pageIndex + 1)
    },
    populatePage(page) {
      this.pageIndex = page
      // eslint-disable-next-line no-undef
      PopulatePage(page, this.pageSize)
          .then(result => this.pageData = result)
    },
    previousPage() {
      this.populatePage(this.pageIndex - 1)
    },
    reset() {
      this.startDate = null
      this.endDate = null

      this.pageData = []
      this.pageIndex = 0

      this.dataSize = 0

      // eslint-disable-next-line no-undef
      Reset()
    }
  }
}
</script>

<style scoped>

</style>