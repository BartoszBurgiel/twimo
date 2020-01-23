import React from "react";
import { Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Define Component

const RatingStars = ({ rating, colorFill, iconSize }) => {
  // Set up targetArray
  let arrayOfIcons = [];
  // Loop ratingValue times
  for (let i = 0; i < 5; i++) {
    if (i < rating) {
      arrayOfIcons.push(
        <Ionicons
          name="md-star"
          key={i}
          size={iconSize ? iconSize : 25}
          color={colorFill}
        />
      );
    } else {
      arrayOfIcons.push(
        <Ionicons
          name="md-star-outline"
          key={i}
          size={iconSize ? iconSize : 25}
          color={colorFill}
        />
      );
    }
  }
  //Return text element containing the array
  return <Text>{arrayOfIcons}</Text>;
};

export default RatingStars;
