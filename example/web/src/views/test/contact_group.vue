<template>
    <div class="app-iframe-container">
        <!-- 搜索 -->
        <el-row class="query-form">
            <el-form label-suffix="：" label-width="150px" label-position="right">
                <el-row>
                    <el-col :span="6">
                        <el-form-item label="Name">
                            <el-input v-model="queryParam.filter.name"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="EmailEnable">
                            <el-input v-model="queryParam.filter.emailEnable"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="DingTalkEnable">
                            <el-input v-model="queryParam.filter.dingTalkEnable"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="WebhookEnable">
                            <el-input v-model="queryParam.filter.webhookEnable"></el-input>
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
                    <el-table-column :show-overflow-tooltip="true" label="webhookEnable" prop="WebhookEnable"/>
                    <el-table-column :show-overflow-tooltip="true" label="dingTalkEnable" prop="DingTalkEnable"/>
                    <el-table-column :show-overflow-tooltip="true" label="emailEnable" prop="EmailEnable"/>
                    <el-table-column :show-overflow-tooltip="true" label="name" prop="Name"/>
                    <el-table-column :show-overflow-tooltip="true" label="updatedAt">
                        <template slot-scope="scope">
                            {{ scope.row.UpdatedAt | parseDateTime }}
                        </template>
                    </el-table-column>
                    <el-table-column :show-overflow-tooltip="true" label="createdAt">
                        <template slot-scope="scope">
                            {{ scope.row.CreatedAt | parseDateTime }}
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
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="name" label="Name">
                    <el-input v-model="editDialog.name"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="emailEnable" label="EmailEnable">
                    <el-input v-model="editDialog.emailEnable"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="emailConfig" label="EmailConfig">
                    <el-input v-model="editDialog.emailConfig"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="dingTalkEnable" label="DingTalkEnable">
                    <el-input v-model="editDialog.dingTalkEnable"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="dingTalkConfig" label="DingTalkConfig">
                    <el-input v-model="editDialog.dingTalkConfig"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="webhookEnable" label="WebhookEnable">
                    <el-input v-model="editDialog.webhookEnable"/>
                </el-form-item>
                <el-form-item :rules="[{required: true, message: '不能为空'}]" prop="webhookConfig" label="WebhookConfig">
                    <el-input v-model="editDialog.webhookConfig"/>
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
                        emailEnable: null,
                        dingTalkEnable: null,
                        webhookEnable: null,
                    }
                },
                editDialog: {
                        show: false,
                        id: null,
                        name: null,
                        emailEnable: null,
                        emailConfig: null,
                        dingTalkEnable: null,
                        dingTalkConfig: null,
                        webhookEnable: null,
                        webhookConfig: null,
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
                this.editDialog.emailEnable = null
                this.editDialog.emailConfig = null
                this.editDialog.dingTalkEnable = null
                this.editDialog.dingTalkConfig = null
                this.editDialog.webhookEnable = null
                this.editDialog.webhookConfig = null
                this.editDialog.show = true
            },
            toEdit() {
                var selectItem = this.multipleSelection[0]
                this.editDialog.id = selectItem.id
                this.editDialog.name = selectItem.name
                this.editDialog.emailEnable = selectItem.emailEnable
                this.editDialog.emailConfig = selectItem.emailConfig
                this.editDialog.dingTalkEnable = selectItem.dingTalkEnable
                this.editDialog.dingTalkConfig = selectItem.dingTalkConfig
                this.editDialog.webhookEnable = selectItem.webhookEnable
                this.editDialog.webhookConfig = selectItem.webhookConfig
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
											query contactGroupPage($data: ContactGroupPageInput!){
												contactGroupPage(data: $data) {
													total
													records {
														id
                						createdAt
                						updatedAt
                						name
                						emailEnable
                						emailConfig
                						dingTalkEnable
                						dingTalkConfig
                						webhookEnable
                						webhookConfig
													}
												}
											}`,
                    variables: {
                        data: this.queryParam
                    },
                    fetchPolicy: 'network-only'
                }).then(data => {
                    this.data = data.data.contactGroupPage.records || []
                    this.total = data.data.contactGroupPage.total || 0
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
                            requestParam.emailEnable = this.editDialog.emailEnable
                            requestParam.emailConfig = this.editDialog.emailConfig
                            requestParam.dingTalkEnable = this.editDialog.dingTalkEnable
                            requestParam.dingTalkConfig = this.editDialog.dingTalkConfig
                            requestParam.webhookEnable = this.editDialog.webhookEnable
                            requestParam.webhookConfig = this.editDialog.webhookConfig
                            this.$apollo.mutate({
                                mutation: gql`mutation editContactGroup($data: ContactGroupEditInput!){
																	editContactGroup(data: $data)
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
                            requestParam.emailEnable = this.editDialog.emailEnable
                            requestParam.emailConfig = this.editDialog.emailConfig
                            requestParam.dingTalkEnable = this.editDialog.dingTalkEnable
                            requestParam.dingTalkConfig = this.editDialog.dingTalkConfig
                            requestParam.webhookEnable = this.editDialog.webhookEnable
                            requestParam.webhookConfig = this.editDialog.webhookConfig
                            this.$apollo.mutate({
                                mutation: gql`mutation addContactGroup($data: ContactGroupAddInput!){
																	addContactGroup(data: $data)
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
                        mutation: gql`mutation removeContactGroups($data: ContactGroupRemovesInput!){
																	removeContactGroups(data: $data)
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