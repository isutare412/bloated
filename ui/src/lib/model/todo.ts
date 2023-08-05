export type UserId = string

export interface Todo {
	id: number
	userId: string
	task: string
	createTime: Date
	updateTime: Date
}

export interface CreateTodoRequest {
	userId: string
	task: string
}

export interface ListTodoReponse {
	todos: Todo[]
}
