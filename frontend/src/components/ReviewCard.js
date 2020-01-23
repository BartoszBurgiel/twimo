// Import core functionality
import React from "react";
import { Text, View, StyleSheet, TouchableOpacity, Image } from "react-native";

// Import color utility
import color from "../../utils/colorPallet";

// Import custom components
import RatingStars from "./RatingStars";

const ReviewCard = ({ reviewData }) => {
  return (
    <TouchableOpacity style={styles.card}>
      <View style={styles.userData}>
        <Image style={styles.image} source={{ uri: reviewData.avatar }}></Image>
        <Text style={styles.userName}>{reviewData.name}</Text>
        <RatingStars
          rating={reviewData.rating}
          colorFill={color.black}
          iconSize={15}
        />
      </View>
      <Text style={styles.reviewContent}>{reviewData.content}</Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: color.white,
    height: 130,
    padding: 10,
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
  userData: {
    flexDirection: "row",
    alignItems: "center"
  },
  userName: {
    fontSize: 15,
    fontWeight: "700",
    marginHorizontal: 8
  },
  image: {
    height: 30,
    width: 30,
    borderRadius: 1061997
  },
  reviewContent: {
    marginTop: 10,
    fontWeight: "200"
  }
});

export default ReviewCard;
