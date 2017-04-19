(function(Vue) {
    "use strict";

    var axios = require('axios');

    new Vue({
        el: '#input_field',
        data: {
            content: ""
        },
        methods: {
            registerTodo: function() {
                axios.post("/api/v1/todo", {
                    text: this.content
                }).then((res) => {
                    console.log("output");
                });
            }
        }
    })

    new Vue({
        el: '#app',
        data: {
            todos: [],
        },
        created: function() {
            axios.get('/api/v1/todos').then((res) => {
                this.todos = res.data;
            }).catch((res) => {
                console.log("error");
            })
        },
        methods: {
            checkedTodo: function(todo_id) {
                console.log(todo_id)
                axios.put(`/api/v1/checked/${todo_id}`).then((res) => {
                    console.log(`checked: ${todo_id}`);
                })
            }
        }
    });
})(Vue);
