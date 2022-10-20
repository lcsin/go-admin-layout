<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <!-- 搜索 -->
        <el-form ref="queryForm" :model="queryParams" :inline="true">
          <el-form-item label="服务名" prop="name">
            <el-input v-model="queryParams.name" placeholder="请输入姓名" clearable size="small" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <!-- 工具栏 -->
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button type="primary" icon="el-icon-plus" size="mini" @click="handleAdd">新增</el-button>
          </el-col>
        </el-row>

        <!-- 表格 -->
        <el-table v-loading="loading" :data="exampleList">
          <el-table-column width="20px" />
          <el-table-column label="ID" prop="id" width="80px" />
          <el-table-column label="姓名" prop="name" :show-overflow-tooltip="true" />
          <el-table-column label="创建时间" :formatter="createTimeFormat" prop="create_time" />
          <el-table-column label="更新时间" :formatter="updateTimeFormat" prop="update_time" />
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
              <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination v-show="total>0" :total="total" :page.sync="queryParams.page" :limit.sync="queryParams.size" @pagination="getList" />

        <!-- 添加或修改对话框 -->
        <el-dialog v-if="open" :title="title" :visible.sync="open" width="50%" :close-on-click-modal="false">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="姓名" prop="name">
              <el-input v-model="form.name" placeholder="请输入姓名" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submit">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  listExample,
  getExample,
  addExample,
  updateExample,
  deleteExample,
} from "@/api/crmexample";

export default {
  name: "CrmExample",
  data() {
    return {
      loading: true,
      queryParams: {
        page: 1,
        size: 10,
        name: "",
      },
      exampleList: [],
      total: 0,
      title: "",
      form: "",
      rules: {
        name: [{ required: true, message: "姓名不能为空", trigger: "blur" }],
      },
      open: false,
    };
  },
  created() {
    this.getList();
  },
  methods: {
    getList() {
      this.loading = true;
      listExample(this.queryParams).then((response) => {
        this.exampleList = response.data.list;
        this.total = response.data.count;
        this.loading = false;
      });
    },
    handleQuery() {
      this.queryParams.page = 1;
      this.getList();
    },
    resetQuery() {
      this.queryParams = {
        page: 1,
        size: 10,
      };
      this.handleQuery();
    },
    handleAdd() {
      this.title = "添加Example";
      this.form = {};
      this.open = true;
    },
    handleUpdate(row) {
      getExample(row.id).then((response) => {
        this.title = "编辑Example";
        this.form = response.data;
        this.open = true;
      });
    },
    handleDelete(row) {
      this.$confirm('是否确认删除ID为"' + row.id + '"的数据项?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(function () {
          return deleteExample(row.id);
        })
        .then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        })
        .catch(function () {});
    },
    // 时间格式化
    createTimeFormat(row) {
      return this.parseTime(row.create_time);
    },
    updateTimeFormat(row) {
      return this.parseTime(row.update_time);
    },
    submit() {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateExample(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess("更新成功");
                this.getList();
                this.open = false;
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addExample(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess("添加成功");
                this.getList();
                this.open = false;
              } else {
                this.msgError(response.msg);
              }
            });
          }
        }
      });
    },
    cancel() {
      this.open = false;
    },
  },
};
</script>

<style scoped>
</style>