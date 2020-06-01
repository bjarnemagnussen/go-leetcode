# 165. Compare Version Numbers

Compare two version numbers `version1` and `version2`.


If `version1 > version2` return `1;` if `version1 < version2` return `-1;`otherwise return `0`.

You may assume that the version strings are non-empty and contain only digits and the `.` character.

The `.` character does not represent a decimal point and is used to separate number sequences.

For instance, `2.5` is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

You may assume the default revision number for each level of a version number to be `0`. For example, version number `3.4` has a revision number of `3` and `4` for its first and second level revision number. Its third and fourth level revision number are both `0`.

## Example 1

<pre>
<b>Input:</b> version1 = "0.1", version2 = "1.1"
<b>Output:</b> -1
</pre>

## Example 2

<pre>
<b>Input:</b> version1 = "1.0.1", version2 = "1"
<b>Output:</b> 1
</pre>

## Example 3

<pre>
<b>Input:</b> version1 = "7.5.2.4", version2 = "7.5.3"
<b>Output:</b> -1
</pre>

## Example 4

<pre>
<b>Input:</b> version1 = "1.01", version2 = "1.001"
<b>Output:</b> 0
<b>Explanation:</b> Ignoring leading zeroes, both “01” and “001" represent the same number “1”
</pre>

## Example 5

<pre>
<b>Input:</b> version1 = "1.0", version2 = "1.0.0"
<b>Output:</b> 0
<b>Explanation:</b> The first version number does not have a third level revision number, which means its third level revision number is default to "0"
</pre> 

</br>

## Note

1. Version strings are composed of numeric strings separated by dots `.` and this numeric strings **may** have leading zeroes.
2. Version strings do not start or end with dots, and they will not be two consecutive dots.
