import React from "react";
import { Text, View, StyleSheet, TouchableOpacity, Image } from "react-native";

import color from "../../utils/colorPallet";

const ReviewCard = () => {
  return (
    <View style={styles.card}>
      <View style={styles.userData}>
        <Image
          style={styles.image}
          source={{ uri: "https://api.adorable.io/avatars/50/someUser.png" }}
        ></Image>
        <Text>UserName</Text>
      </View>
      <Text>ReviewContent</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: color.white,
    height: 125,
    padding: 15,
    marginVertical: 10,
    shadowColor: color.gray.darker,
    shadowOpacity: 0.25,
    borderRadius: 12,
    shadowOffset: {
      width: 0,
      height: 3
    },
    shadowRadius: 13,
    flexDirection: "column"
  },
  userData: {},
  image: {
    height: 30,
    width: 30,
    borderRadius: 1061997
  }
});

export default ReviewCard;
