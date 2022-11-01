import {View, Text, StyleSheet} from 'react-native'

function Board_item({id, author, title, content, diffMin}) {
    return (
        <View style={styles.container}>
            <View>
                <Text style={styles.title}>{title}</Text>
            </View>
            <View>
                <Text style={styles.content}>{content}</Text>
            </View>
            <View>
                <Text style={styles.time_name}>{diffMin} | {author}</Text>
            </View>
        </View>
    )
}

export default Board_item

const styles = StyleSheet.create({
    container: {
        marginBottom: "4%",
        borderBottomWidth: 1,
        borderBottomColor: "rgba(0, 0, 0, 0.2)",
        paddingBottom: "2%"
    },
    
    title: {
        fontSize: 18
    },

    content: {
        fontSize: 14,
    },

    time_name: {
        fontSize: 11,
        color: "rgba(0, 0, 0, 0.6)"
    }

})