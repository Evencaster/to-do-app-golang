import {
    atom
} from "recoil"

export const tasks = atom({
    key: 'tasks',
    default: [],
})

export const newTaskName = atom({
    key: 'newTaskName',
    default: [],
})