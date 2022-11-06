<template>
  <div>
    <el-row>
      选择需要生成的日期
    </el-row>
    <el-row>
      <el-col :span="10">
        <el-date-picker v-model="time" type="daterange" align="right" format="yyyy-MM-dd" value-format="timestamp"
          unlink-panels range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期"
          :picker-options="pickerOptions" @change="dateChange">
        </el-date-picker>
      </el-col>
    </el-row>
    <el-row>
      日期之间分割字段（输入\n可以换行）
    </el-row>
    <el-row>
      <el-col>
        <el-input v-model="input" placeholder="请输入内容" @input="dateChange"></el-input>
      </el-col>
    </el-row>
    <el-row>
      处理完的字段
    </el-row>
    <el-row>
      <el-col>
        <el-input type="textarea" :autosize="{ minRows: 10, maxRows: 20}" placeholder="生成的内容" v-model="textarea">
        </el-input>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        pickerOptions: {
          shortcuts: [{
            text: '最近一周',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit('pick', [start, end]);
            }
          }, ],
        },
        time: [Date.now(), Date.now() + 60 * 60 * 24 * 1000],
        input: ",",
        textarea: "",
      };
    },
    created() {
      this.dateChange()
    },
    methods: {
      async dateChange() {
        if (!this.time) {
          return
        }
        const {
          data: res
        } = await this.$http.post(`dt_deal`, {
          start_unix: parseInt(this.time[0] / 1000),
          end_unix: parseInt(this.time[1] / 1000),
          glue: this.input
        })
        if (res.status != 200) {
          this.$message.error(res.message)
        } else {
          this.textarea = res.data
        }
      },
    }
  };
</script>

<style>
  .el-row {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .el-col {
    border-radius: 4px;
  }

  .bg-purple-dark {
    background: #99a9bf;
  }

  .bg-purple {
    background: #d3dce6;
  }

  .bg-purple-light {
    background: #e5e9f2;
  }

  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
</style>
¶