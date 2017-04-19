var axios = require('axios');

var getTodos = () => {
    return axios.get('/api/v1/todos')
}

(function(Vue) {
    "use strict";

    new Vue({
        el: '#app',
        data: {
            todos: [],
            content: "",
        },
        created: function() {
            getTodos().then((res) => {
                this.todos = res.data;
            }).catch((res) => {
                console.log("error");
            });
        },
        methods: {
            registerTodo: function() {
                axios.post("/api/v1/todo", {
                    text: this.content,
                }).then((res) => {
                    getTodos().then((res) => {
                        this.todos = res.data;
                    });
                });
            },
            checkedTodo: function(todo_id) {
                console.log(todo_id)
                axios.put(`/api/v1/checked/${todo_id}`).then((res) => {
                    console.log(`checked: ${todo_id}`);
                })
            }
        }
    });
})(Vue);
