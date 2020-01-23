import React from "react";
import { Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Define Component

const RatingCounter = ({ rating, colorfill }) => {
  // Set up targetArray
  let arrayOfIcons = [];
  // Loop ratingValue times
  for (let i = 0; i < 5; i++) {
    if (i < rating) {
      arrayOfIcons.push(
        <Ionicons name="md-star" key={i} size={25} color={colorfill} />
      );
    } else {
      arrayOfIcons.push(
        <Ionicons name="md-star-outline" key={i} size={25} color={colorfill} />
      );
    }
  }
  //Return text element containing the array
  return <Text>{arrayOfIcons}</Text>;
};

export default RatingCounter;
