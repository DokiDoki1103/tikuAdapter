export const questionType = {
    '0': '单选题',
    '1': '多选题',
    '2': '填空题',
    '3': '判断题',
    '4': '简答题',
    '-4': '大作业'
}

export function getQuestionTypeByName(name) {
    for (let key of Object.keys(questionType)) {
        if (name === questionType[key]) {
            return parseInt(key);
        }
    }
    return 4
}