<template>
    <div class="app-iframe-container">
        <!-- 搜索 -->
        <el-row class="query-form">
            <el-form label-suffix="：" label-width="150px" label-position="right">
                <el-row>
                    <el-col :span="6">
                        <el-form-item label="名称">
                            <el-input v-model="queryParam.filter.name"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="邮箱">
                            <el-input v-model="queryParam.filter.email"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="手机号码">
                            <el-input v-model="queryParam.filter.phone"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="24" el-col style="text-align: center;">
                        <el-button type="warning" size="small" icon="el-icon-search" @click="research">查询</el-button>
                    </el-col>
                </el-row>
            </el-form>
        </el-row>
        <!-- /搜索 -->
        <el-row class="button-row">
            <el-col :span="24">
                <el-button size="small" @click="toAdd()">添加</el-button>
                <el-button :disabled="!editButtonValid" size="small" type="primary" @click="toEdit()">修改</el-button>
                <el-button :disabled="!removesButtonValid" size="small" type="danger" @click="removes()">删除</el-button>
            </el-col>
        </el-row>
        <!-- /操作按钮 -->
        <!-- 表格内容 -->
        <el-row>
            <el-col :span="24">
                <el-table
                        :data="data"
                        :loading="listLoading"
                        border
                        fit
                        highlight-current-row
                        style="width: 100%;"
                        @selection-change="handleSelectionChange">
                    <el-table-column
                            type="selection"
                            width="55"/>
                    <el-table-column align="center" width="65" label="序列号">
                        <template slot-scope="scope">
                            {{ scope.$index + 1 }}
                        </template>
                    </el-table-column>
                    <el-table-column :show-overflow-tooltip="true" label="手机号码" prop="phone"/>
                    <el-table-column :show-overflow-tooltip="true" label="邮箱" prop="email"/>
                    <el-table-column :show-overflow-tooltip="true" label="名称" prop="name"/>
                    <el-table-column :show-overflow-tooltip="true" label="更新时间">
                        <template slot-scope="scope">
                            {{ scope.row.updatedAt | parseDateTime }}
                        </template>
                    </el-table-column>
                    <el-table-column :show-overflow-tooltip="true" label="创建时间">
                        <template slot-scope="scope">
                            {{ scope.row.createdAt | parseDateTime }}
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
        </el-row>
        <!-- /表格内容 -->
        <!-- 分页 -->
        <div class="pagination-container">
            <el-pagination v-show="total>0" :current-page="queryParam.pageNum" :page-sizes="appCommon.pageSize" :page-size="queryParam.pageSize" :total="total" background layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
        </div>
        <!-- /分页 -->
        <el-dialog :visible.sync="editDialog.show">
            <el-form ref="editDialog" :model="editDialog">
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="name" label="名称">
                    <el-input v-model="editDialog.name"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="email" label="邮箱">
                    <el-input v-model="editDialog.email"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="phone" label="手机号码">
                    <el-input v-model="editDialog.phone"/>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="doEdit">保存</el-button>
                    <el-button @click="editDialog.show = false">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>
<script>
    import gql from 'graphql-tag'
    export default {
        name: 'DcronService',
        data() {
            return {
                listLoading: false,
                multipleSelection: [],
                data: [],
                total: 0,
                queryParam: {
                    pageNum: 1,
                    pageSize: 10,
                    filter: {
                        name: null,
                        email: null,
                        phone: null,
                    }
                },
                editDialog: {
                        show: false,
                        id: null,
                        name: null,
                        email: null,
                        phone: null,
                },
            }
        },
        computed: {
            removesButtonValid() {
                if (!this.multipleSelection || this.multipleSelection.length === 0) {
                    return false
                }
                return true
            },
            editButtonValid() {
                if (!this.multipleSelection || this.multipleSelection.length === 0) {
                    return false
                }
                if (this.multipleSelection.length === 1) {
                    return true
                }
                return false
            }
        },
        mounted() {
            this.research()
        },
        methods: {
            handleSizeChange(val) {
                this.queryParam.pageSize = val
                this.research()
            },
            // 处理页数改变
            handleCurrentChange(val) {
                this.queryParam.pageNum = val
                this.search()
            },
            toAdd() {
                this.editDialog.id = null
                this.editDialog.name = null
                this.editDialog.email = null
                this.editDialog.phone = null
                this.editDialog.show = true
            },
            toEdit() {
                var selectItem = this.multipleSelection[0]
                this.editDialog.id = selectItem.id
                this.editDialog.name = selectItem.name
                this.editDialog.email = selectItem.email
                this.editDialog.phone = selectItem.phone
                this.editDialog.show = true
            },
            // 处理选择的变化
            handleSelectionChange(val) {
                this.multipleSelection = val
            },
            // 列表
            search() {
                this.listLoading = true
                this.$apollo.query({
                    query: gql`
											query contactPage($data: ContactPageInput!){
												contactPage(data: $data) {
													total
													records {
														id
                						createdAt
                						updatedAt
                						name
                						email
                						phone
													}
												}
											}`,
                    variables: {
                        data: this.queryParam
                    },
                    fetchPolicy: 'network-only'
                }).then(data => {
                    this.data = data.data.contactPage.records || []
                    this.total = data.data.contactPage.total || 0
                }).finally(()=>{
                    this.listLoading = false
                })
            },
            // 重置分页条件，发起查询
            research() {
                this.queryParam.pageNum = 1
                this.search()
            },
            doEdit() {
                this.$refs.editDialog.validate((valid) => {
                    if (valid) {
                        if (this.editDialog.id) {
                            // 修改
                            const requestParam = {}
                            requestParam.id = this.editDialog.id
                            requestParam.name = this.editDialog.name
                            requestParam.email = this.editDialog.email
                            requestParam.phone = this.editDialog.phone
                            this.$apollo.mutate({
                                mutation: gql`mutation editContact($data: ContactEditInput!){
																	editContact(data: $data)
															}`,
                                variables: {
                                    data: requestParam
                                }
                            }).then(data => {
                                this.editDialog.show = false
                                this.research()
                                this.$notify.success('操作成功！') // 提示
                            }).catch(err => {
                                this.$notify.error(err.message)
                            })
                        } else {
                            // 新增
                            const requestParam = {}
                            requestParam.name = this.editDialog.name
                            requestParam.email = this.editDialog.email
                            requestParam.phone = this.editDialog.phone
                            this.$apollo.mutate({
                                mutation: gql`mutation addContact($data: ContactAddInput!){
																	addContact(data: $data)
															}`,
                                variables: {
                                    data: requestParam
                                }
                            }).then(data => {
                                this.editDialog.show = false
                                this.research()
                                this.$notify.success('操作成功！') // 提示
                            }).catch(err => {
                                this.$notify.error(err.message)
                            })
                        }
                    }
                })
            },
            removes() {
                this.$confirm('确定删除吗', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    var requestParam = {}
                    requestParam.ids = []
                    for (const item of this.multipleSelection) {
                        requestParam.ids.push(item.id)
                    }
                    this.$apollo.mutate({
                        mutation: gql`mutation removeContacts($data: ContactRemovesInput!){
																	removeContacts(data: $data)
															}`,
                        variables: {
                            data: requestParam
                        }
                    }).then(data => {
                        this.research()
                        this.$notify.success('操作成功！') // 提示
                    }).catch(err => {
                        this.$notify.error(err.message)
                    })
                })
            }
        }
    }
</script>
<style scoped>
    .el-row {
        margin-bottom: 20px;
    }
    .danger {
        color: #F56C6C
    }
</style>